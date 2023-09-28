package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	_ "github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image/color"
	"log"
	"os"
)

type CPU struct {
	memory [4096]byte
	v      [16]uint8
	pc     uint16
	i      uint16
	sp     uint16
	stack  [16]uint16
}

type Chip8 struct {
	cpu     CPU
	clavier Clavier
	screen  Screen
}

type Clavier struct {
	Key []string
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
			fmt.Println("CLS")
			for i := 0; i < resolWidth; i++ {
				for j := 0; j < resolHeight; j++ {
					chip8.screen.mapscreen[i][j] = 0
				}
			}
		case 0x000E:
			fmt.Println("RET")
			//0x000E RET -> Return from a subroutine.
			chip8.cpu.pc = uint16(int(chip8.cpu.memory[chip8.cpu.pc-1]))
		}
	case 0x1000:
		fmt.Printf("JP addr = %x\n", b&0x0FFF)
		//0x1NNN JP addr -> Jump to location nnn.
		chip8.cpu.pc = b&0x0FFF - 2
	case 0x2000:
		fmt.Println("CALL addr")
		//0x2NNN CALL addr -> Call subroutine at nnn.
		chip8.cpu.stack[chip8.cpu.sp] = chip8.cpu.pc
		chip8.cpu.sp++
		chip8.cpu.pc += 2
	case 0x3000:
		fmt.Println("SE Vx, byte")
		//0x3XNN SE Vx, byte -> Skip next instruction if Vx = kk.
		if chip8.cpu.v[(b&0x0F00)>>8] == uint8(b&0x00FF) {
			chip8.cpu.pc += 2
		}
	case 0x4000:
		fmt.Println("SNE Vx, byte")
		//0x4XNN SNE Vx, byte -> Skip next instruction if Vx != kk.
		if chip8.cpu.v[(b&0x0F00)>>8] != uint8(b&0x00FF) {
			chip8.cpu.pc += 2
		}
	case 0x5000:
		fmt.Println(" SE Vx, Vy")
		//0x5XY0 SE Vx, Vy -> Skip next instruction if Vx = Vy.
		if chip8.cpu.v[(b&0x0F00)>>8] == chip8.cpu.v[(b&0x00F0)>>4] {
			chip8.cpu.pc += 2
		}
	case 0x6000:
		fmt.Println("LD Vx, byte")
		//0x6XNN LD Vx, byte -> Set Vx = kk.
		chip8.cpu.v[(b&0x0F00)>>8] = uint8(b & 0x00FF)
	case 0x7000:
		fmt.Println("ADD Vx, byte")
		//0x7XNN ADD Vx, byte -> Set Vx = Vx + kk.
		chip8.cpu.v[(b&0x0F00)>>8] += uint8(b & 0x00FF)
	case 0x8000:
		switch b & 0x000F {
		case 0x0000:
			fmt.Println("LD Vx, Vy")
		case 0x0001:
			fmt.Println("OR Vx, Vy")
		case 0x0002:
			fmt.Println("AND Vx, Vy")
		case 0x0003:
			fmt.Println("XOR Vx, Vy")
		case 0x0004:
			fmt.Println("ADD Vx, Vy")
		case 0x0005:
			fmt.Println("SUB Vx, Vy")
		case 0x0006:
			fmt.Println("SHR Vx {, Vy}")
		case 0x0007:
			fmt.Println("SUBN Vx, Vy")
		case 0x000E:
			fmt.Println("SHL Vx {, Vy}")
		}
	case 0x9000:
		fmt.Println("SNE Vx, Vy")
	case 0xA000:
		fmt.Println("LD I, addr")
	case 0xB000:
		fmt.Println("JP V0, addr")
	case 0xC000:
		fmt.Println("RND Vx, byte")
	case 0xD000:
		fmt.Println("DRW Vx, Vy, nibble")
		var tmps []uint8
		for i := chip8.cpu.i; i < b&0x000F; i++ {
			tmps = append(tmps, chip8.cpu.memory[i])
		}
		for i := uint8(0); i < uint8(len(tmps)); i++ {
			for j := uint8(0); j < 8; j++ {
				if tmps[i]&(0x80>>j) != 0 {
					chip8.screen.mapscreen[chip8.cpu.v[(b&0x0F00)>>8]+i][chip8.cpu.v[(b&0x00F0)>>4]+j] ^= 1
				}
			}
		}
	case 0xE000:
		switch b & 0x000F {
		case 0x000E:
			fmt.Println("SKP Vx")
		case 0x0001:
			fmt.Println("SKNP Vx")
		}
	case 0xF000:
		switch b & 0x000F {
		case 0x0007:
			fmt.Println("Fx07 - LD Vx, DT")
		case 0x000A:
			fmt.Println("LD Vx, K")
		case 0x0005:
			switch b & 0x00F0 {
			case 0x0010:
				fmt.Println("LD DT, Vx")
			case 0x0050:
				fmt.Println("LD [I], Vx")
			case 0x0060:
				fmt.Println("LD Vx, [I]")
			}
		case 0x0008:
			fmt.Println("LD ST, Vx")
		case 0x000E:
			fmt.Println("ADD I, Vx")
		case 0x0009:
			fmt.Println("LD F, Vx")
		case 0x0003:
			fmt.Println("LD B, Vx")
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

func unit16to8(a uint16) (uint8, uint8) {
	return uint8(a >> 8), uint8(a & 0x00FF)
}

func unit8to4(a uint8) (uint8, uint8) {
	return uint8(a >> 4), uint8(a & 0x0F)
}
