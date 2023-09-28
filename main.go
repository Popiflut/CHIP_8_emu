package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	_ "github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image/color"
	"log"
	"math/rand"
	"os"
)

type CPU struct {
	memory [4096]byte
	v      [16]uint8
	pc     uint16
	i      uint16
	dt     uint8
	st     uint8
	sp     uint16
	stack  [16]uint16
}

type Chip8 struct {
	cpu     CPU
	clavier Clavier
	screen  Screen
}

type Clavier struct {
	isPressed [16]bool
}

type Screen struct {
	mapscreen [64][32]uint8
}

var (
	chip8  = Chip8{}
	ROM    []byte
	screen *ebiten.Image
)

const (
	screenWidth  = 640
	screenHeight = 320
	resolWidth   = 64
	resolHeight  = 32
)

//----------------------------------------------------------------------------------------                                                    FRONT

// Console représente l'état du jeu.
type Console struct {
	IN      string
	OUT     string
	command string
}

func RefreshKeyBoard() {
	chip8.clavier.isPressed = [16]bool{
		ebiten.IsKeyPressed(ebiten.Key1),
		ebiten.IsKeyPressed(ebiten.Key2),
		ebiten.IsKeyPressed(ebiten.Key3),
		ebiten.IsKeyPressed(ebiten.Key4),
		ebiten.IsKeyPressed(ebiten.KeyA),
		ebiten.IsKeyPressed(ebiten.KeyZ),
		ebiten.IsKeyPressed(ebiten.KeyE),
		ebiten.IsKeyPressed(ebiten.KeyR),
		ebiten.IsKeyPressed(ebiten.KeyQ),
		ebiten.IsKeyPressed(ebiten.KeyS),
		ebiten.IsKeyPressed(ebiten.KeyD),
		ebiten.IsKeyPressed(ebiten.KeyF),
		ebiten.IsKeyPressed(ebiten.KeyW),
		ebiten.IsKeyPressed(ebiten.KeyX),
		ebiten.IsKeyPressed(ebiten.KeyC),
		ebiten.IsKeyPressed(ebiten.KeyV),
	}
}

func Init() {
	var fontSet = []byte{
		0xF0, 0x90, 0x90, 0x90, 0xF0, // 0
		0x20, 0x60, 0x20, 0x20, 0x70, // 1
		0xF0, 0x10, 0xF0, 0x80, 0xF0, // 2
		0xF0, 0x10, 0xF0, 0x10, 0xF0, // 3
		0x90, 0x90, 0xF0, 0x10, 0x10, // 4
		0xF0, 0x80, 0xF0, 0x10, 0xF0, // 5
		0xF0, 0x80, 0xF0, 0x90, 0xF0, // 6
		0xF0, 0x10, 0x20, 0x40, 0x40, // 7
		0xF0, 0x90, 0xF0, 0x90, 0xF0, // 8
		0xF0, 0x90, 0xF0, 0x10, 0xF0, // 9
		0xF0, 0x90, 0xF0, 0x90, 0x90, // A
		0xE0, 0x90, 0xE0, 0x90, 0xE0, // B
		0xF0, 0x80, 0x80, 0x80, 0xF0, // C
		0xE0, 0x90, 0x90, 0x90, 0xE0, // D
		0xF0, 0x80, 0xF0, 0x80, 0xF0, // E
		0xF0, 0x80, 0xF0, 0x80, 0x80, // F
	}
	for i := 0; i < len(fontSet); i++ {
		chip8.cpu.memory[i] = fontSet[i]
	}
	for i := 0; i < len(chip8.clavier.isPressed); i++ {
		chip8.clavier.isPressed[i] = false
	}
}

func (kb *Clavier) GetKey(key byte) bool {
	return kb.isPressed[key]
}

// NewConsole initialise un nouveau jeu.
func NewConsole() *Console {
	console := &Console{}
	return console
}

func (g *Console) Update() error {
	chip8.cpu.pc += 2
	fmt.Printf("cp = %02X:0x%04X: ", chip8.cpu.pc, (uint16(chip8.cpu.memory[chip8.cpu.pc])<<8)|uint16(chip8.cpu.memory[chip8.cpu.pc+1]))
	chip8.cpu.Interpreter((uint16(chip8.cpu.memory[chip8.cpu.pc]) << 8) | uint16(chip8.cpu.memory[chip8.cpu.pc+1]))
	RefreshKeyBoard()
	return nil
}

func (g *Console) Draw(screen *ebiten.Image) {
	for x, row := range chip8.screen.mapscreen {
		for y, pixel := range row {
			var c color.Color
			if pixel == 1 {
				c = color.White
			} else {
				c = color.Black
			}
			screen.Set(x, y, c)
		}
	}
}

func (g *Console) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return resolWidth, resolHeight
}

func LoadProgram() {
	//s = new
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("CHIP-8 Console")
	console := NewConsole()
	ebiten.SetTPS(60)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	if err := ebiten.RunGame(console); err != nil {
		log.Fatal(err)
	}
}

//----------------------------------------------------------------------------------------                                                    FRONT END

func LoadROM(file []byte) {
	chip8.cpu.memory = [4096]byte{}
	n := 0x200
	for i := 0; i < len(file); i++ {
		chip8.cpu.memory[n] = file[i]
		n++
	}
}

func (cpu *CPU) Interpreter(b uint16) {
	switch b & 0xF000 {
	case 0x0000:
		switch b & 0x000F {
		case 0x0000:
			//0x0000 CLS -> Clear the display.
			for i := 0; i < resolWidth; i++ {
				for j := 0; j < resolHeight; j++ {
					chip8.screen.mapscreen[i][j] = 0
				}
			}
		case 0x000E:
			//0x000E RET -> Return from a subroutine.
			chip8.cpu.pc = StackPop()
		}
	case 0x1000:
		//0x1NNN JP addr -> Jump to location nnn.
		chip8.cpu.pc = b&0x0FFF - 2
	case 0x2000:
		//0x2NNN CALL addr -> Call subroutine at nnn.
		StackPush(chip8.cpu.pc)
		chip8.cpu.pc = b&0x0FFF - 2
	case 0x3000:
		//0x3XNN SE Vx, byte -> Skip next instruction if Vx = kk.
		if chip8.cpu.v[(b&0x0F00)>>8] == uint8(b&0x00FF) {
			chip8.cpu.pc += 2
		}
	case 0x4000:
		//0x4XNN SNE Vx, byte -> Skip next instruction if Vx != kk.
		if chip8.cpu.v[(b&0x0F00)>>8] != uint8(b&0x00FF) {
			chip8.cpu.pc += 2
		}
	case 0x5000:
		//0x5XY0 SE Vx, Vy -> Skip next instruction if Vx = Vy.
		if chip8.cpu.v[(b&0x0F00)>>8] == chip8.cpu.v[(b&0x00F0)>>4] {
			chip8.cpu.pc += 2
		}
	case 0x6000:
		//0x6XNN LD Vx, byte -> Set Vx = kk.
		chip8.cpu.v[(b&0x0F00)>>8] = uint8(b & 0x00FF)
	case 0x7000:
		//0x7XNN ADD Vx, byte -> Set Vx = Vx + kk.
		chip8.cpu.v[(b&0x0F00)>>8] += uint8(b & 0x00FF)
	case 0x8000:
		switch b & 0x000F {
		case 0x0000:
			//0x8XY0 LD Vx, Vy -> Set Vx = Vy.
			chip8.cpu.v[(b&0x0F00)>>8] = chip8.cpu.v[(b&0x00F0)>>4]
		case 0x0001:
			//0x8XY1 OR Vx, Vy -> Set Vx = Vx OR Vy.
			chip8.cpu.v[(b&0x0F00)>>8] |= chip8.cpu.v[(b&0x00F0)>>4]
		case 0x0002:
			//0x8XY2 AND Vx, Vy -> Set Vx = Vx AND Vy.
			chip8.cpu.v[(b&0x0F00)>>8] &= chip8.cpu.v[(b&0x00F0)>>4]
		case 0x0003:
			//0x8XY3 XOR Vx, Vy -> Set Vx = Vx XOR Vy.
			chip8.cpu.v[(b&0x0F00)>>8] ^= chip8.cpu.v[(b&0x00F0)>>4]
		case 0x0004:
			//0x8XY4 ADD Vx, Vy -> Set Vx = Vx + Vy, set VF = carry.
			if chip8.cpu.v[(b&0x00F0)>>4] > (0xFF - chip8.cpu.v[(b&0x0F00)>>8]) {
				chip8.cpu.v[0xF] = 1
			} else {
				chip8.cpu.v[0xF] = 0
			}
			chip8.cpu.v[(b&0x0F00)>>8] += chip8.cpu.v[(b&0x00F0)>>4]
		case 0x0005:
			//0x8XY5 SUB Vx, Vy -> Set Vx = Vx - Vy, set VF = NOT borrow.
			if chip8.cpu.v[(b&0x00F0)>>4] > chip8.cpu.v[(b&0x0F00)>>8] {
				chip8.cpu.v[0xF] = 0
			} else {
				chip8.cpu.v[0xF] = 1
			}
			chip8.cpu.v[(b&0x0F00)>>8] -= chip8.cpu.v[(b&0x00F0)>>4]
		case 0x0006:
			//0x8XY6 SHR Vx {, Vy} -> Set Vx = Vx SHR 1.
			chip8.cpu.v[0xF] = chip8.cpu.v[(b&0x0F00)>>8] & 0x1
			chip8.cpu.v[(b&0x0F00)>>8] >>= 1
		case 0x0007:
			//0x8XY7 SUBN Vx, Vy -> Set Vx = Vy - Vx, set VF = NOT borrow.
			if chip8.cpu.v[(b&0x0F00)>>8] > chip8.cpu.v[(b&0x00F0)>>4] {
				chip8.cpu.v[0xF] = 0
			} else {
				chip8.cpu.v[0xF] = 1
			}
			chip8.cpu.v[(b&0x0F00)>>8] = chip8.cpu.v[(b&0x00F0)>>4] - chip8.cpu.v[(b&0x0F00)>>8]
		case 0x000E:
			//0x8XYE SHL Vx {, Vy} -> Set Vx = Vx SHL 1.
			chip8.cpu.v[0xF] = chip8.cpu.v[(b&0x0F00)>>8] >> 7
			chip8.cpu.v[(b&0x0F00)>>8] <<= 1
		}
	case 0x9000:
		//0x9XY0 SNE Vx, Vy -> Skip next instruction if Vx != Vy.
		if chip8.cpu.v[(b&0x0F00)>>8] != chip8.cpu.v[(b&0x00F0)>>4] {
			chip8.cpu.pc += 2
		}
	case 0xA000:
		//0xANNN LD I, addr -> Set I = nnn.
		chip8.cpu.i = b & 0x0FFF
	case 0xB000:
		//0xBNNN JP V0, addr -> Jump to location nnn + V0.
		chip8.cpu.pc = (b & 0x0FFF) + uint16(chip8.cpu.v[0])
	case 0xC000:
		//0xCXNN RND Vx, byte -> Set Vx = random byte AND nn.
		chip8.cpu.v[(b&0x0F00)>>8] = uint8(rand.Intn(255)) & uint8(b&0x00FF)
	case 0xD000:
		//0xDXYK
		var tmps []uint8
		for i := chip8.cpu.i; i < chip8.cpu.i+b&0x000F; i++ {
			tmps = append(tmps, chip8.cpu.memory[i])
		}
		for i := uint8(0); i < uint8(len(tmps)); i++ {
			for j := uint8(0); j < 8; j++ {
				chip8.screen.mapscreen[(chip8.cpu.v[(b&0x0F00)>>8]+j)%64][(chip8.cpu.v[(b&0x00F0)>>4]+i)%32] ^= tmps[i] >> (7 - j) & 0x01
			}
		}
	case 0xE000:
		switch b & 0x000F {
		case 0x000E:
			fmt.Println("\nSKNP Vx ----------------------------------------------------------------------------- Not Pressed")
			//0xEX9E SKP Vx -> Skip next instruction if key with the value of Vx is pressed.
			if chip8.clavier.isPressed[chip8.cpu.v[(b&0x0F00)>>8]] {
				fmt.Println("\nSKP Vx ----------------------------------------------------------------------------- Pressed")
				chip8.cpu.pc += 2
			}
		case 0x0001:

			//0xEXA1 SKNP Vx -> Skip next instruction if key with the value of Vx is not pressed.
			if !chip8.clavier.isPressed[chip8.cpu.v[(b&0x0F00)>>8]] {
				chip8.cpu.pc += 2
			}
		}
	case 0xF000:
		switch b & 0x000F {
		case 0x0007:
			//0xFX07 LD Vx, DT -> Set Vx = delay timer value.
			chip8.cpu.v[(b&0x0F00)>>8] = chip8.cpu.dt
		case 0x000A:
			//0xFX0A LD Vx, K -> Wait for a key press, store the value of the key in Vx.
			chip8.cpu.v[(b&0x0F00)>>8] = chip8.cpu.dt
		case 0x0005:
			switch b & 0x00F0 {
			case 0x0010:
				//0xFX15 LD DT, Vx -> Set delay timer = Vx.
				chip8.cpu.dt = chip8.cpu.v[(b&0x0F00)>>8]
			case 0x0050:
				//0xFX55 LD [I], Vx -> Store registers V0 through Vx in memory starting at location I.
				for i := uint16(0); i <= (b&0x0F00)>>8; i++ {
					chip8.cpu.memory[chip8.cpu.i+i] = chip8.cpu.v[i]
				}
			case 0x0060:
				//0xFX65 LD Vx, [I] -> Read registers V0 through Vx from memory starting at location I.
				for i := uint16(0); i <= (b&0x0F00)>>8; i++ {
					chip8.cpu.v[i] = chip8.cpu.memory[chip8.cpu.i+i]
				}
			}
		case 0x0008:
			//0xFX18 LD ST, Vx -> Set sound timer = Vx.
			chip8.cpu.st = chip8.cpu.v[(b&0x0F00)>>8]
		case 0x000E:
			//0xFX1E ADD I, Vx -> Set I = I + Vx.
			chip8.cpu.i += uint16(chip8.cpu.v[(b&0x0F00)>>8])
		case 0x0009:
			//0xFX29 LD F, Vx -> Set I = location of sprite for digit Vx.
			chip8.cpu.i = uint16(chip8.cpu.v[(b&0x0F00)>>8] * 5)
		case 0x0003:
			//0xFX33 LD B, Vx -> Store BCD representation of Vx in memory locations I, I+1, and I+2.
			chip8.cpu.memory[chip8.cpu.i] = chip8.cpu.v[(b&0x0F00)>>8] / 100
			chip8.cpu.memory[chip8.cpu.i+1] = (chip8.cpu.v[(b&0x0F00)>>8] / 10) % 10
			chip8.cpu.memory[chip8.cpu.i+2] = (chip8.cpu.v[(b&0x0F00)>>8] % 100) % 10
		}
	}
}

func Start() error {
	file, err := os.ReadFile(os.Args[1])
	if err != nil {
		return err
	}
	ROM = file
	Init()
	fmt.Println("Init...")
	//fmt.Println(chip8.clavier.Keyboard)
	chip8.screen.mapscreen = [64][32]uint8{}
	LoadROM(file)
	chip8.cpu.pc = 0x1FE
	ebiten.SetTPS(60)
	fmt.Println("Loading ROM...")
	fmt.Println("ROM size: ", len(file))
	fmt.Println("Chip8 Emulator")
	fmt.Println("ROM: ", os.Args[1])
	fmt.Println("Start")
	LoadProgram()
	return nil
}

func main() {
	err := Start()
	if err != nil {
		fmt.Println("ERROR system start")
		return
	}
}

func StackPush(pc uint16) {
	chip8.cpu.stack[chip8.cpu.sp] = pc
	chip8.cpu.sp++
}

func StackPop() uint16 {
	chip8.cpu.sp--
	return chip8.cpu.stack[chip8.cpu.sp]
}
