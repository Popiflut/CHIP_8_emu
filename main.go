package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
	"log"
	"os"
)

type CPU struct {
	memory [4096]byte
	v      [16]uint8
	pc     int
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
	s [64][32]color.Color
}

var (
	chip8 = Chip8{}
	ROM   = []byte{}
)

const (
	screenWidth  = 640
	screenHeight = 480
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

// Update met à jour l'état du jeu à chaque trame.
func (g *Console) Update() error {
	chip8.cpu.pc += 2
	fmt.Printf("cp = %02X:0x%04X: ", chip8.cpu.pc, (uint16(chip8.cpu.memory[chip8.cpu.pc])<<8)|uint16(chip8.cpu.memory[chip8.cpu.pc+1]))
	chip8.cpu.Interpreter((uint16(chip8.cpu.memory[chip8.cpu.pc]) << 8) | uint16(chip8.cpu.memory[chip8.cpu.pc+1]))
	if chip8.cpu.pc >= len(ROM)+0x200 { // remettre a 0x200
		os.Exit(0)
	}
	return nil
}

// Draw dessine le jeu sur l'écran
func (g *Console) Draw(screen *ebiten.Image) {
	for x := 0; x < len(chip8.screen.s); x++ {
		for y := 0; y < len(chip8.screen.s[x]); y++ {
			//ebitenutil.DrawRect(screen, float64(x*(screenWidth/64)), float64(y*(screenHeight/32)), float64((x+1)*(screenWidth/64)), float64((y+1)*(screenHeight/32)), chip8.screen.s[x][y])
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

func (cpu CPU) Interpreter(b uint16) {
	switch b & 0xF000 {
	case 0x0000:
		switch b & 0x000F {
		case 0x0000:
			//0x0000 CLS -> Clear the display.
			fmt.Println("CLS")
			for i, x := range chip8.screen.s {
				for j, _ := range x {
					chip8.screen.s[i][j] = color.White
				}
			}
		case 0x000E:
			fmt.Println("RET")
			//0x000E RET -> Return from a subroutine.
		}
	case 0x1000:
		fmt.Printf("JP addr = %d\n", int(b&0x0FFF))
		//0x1NNN JP addr -> Jump to location nnn.
		chip8.cpu.pc = int(b & 0x0FFF)
	case 0x2000:
		fmt.Println("CALL addr")
		//0x2NNN CALL addr -> Call subroutine at nnn.
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
			fmt.Println(" LD Vx, Vy")
		case 0x0001:
			fmt.Println("OR Vx, Vy")
		case 0x0002:
			fmt.Println("AND Vx, Vy")
		case 0x0003:
			fmt.Println("XOR Vx, Vy")
		case 0x0004:
			fmt.Println("ADD Vx, Vy")
		case 0x0005:
			fmt.Println(" SUB Vx, Vy")
		case 0x0006:
			fmt.Println("SHR Vx {, Vy}")
		case 0x0007:
			fmt.Println("SUBN Vx, Vy")
		case 0x000E:
			fmt.Println(" SHL Vx {, Vy}")
		}
	case 0x9000:
		fmt.Println("SNE Vx, Vy")
	case 0xA000:
		fmt.Println("LD I, addr")
	case 0xB000:
		fmt.Println(" JP V0, addr")
	case 0xC000:
		fmt.Println("RND Vx, byte")
	case 0xD000:
		fmt.Println(" DRW Vx, Vy, nibble")
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
	return 0
}

func Start() error {
	file, err := os.ReadFile(os.Args[1])
	if err != nil {
		return err
	}
	ROM = file
	Init()
	LoadROM(file)
	chip8.cpu.pc = 0x1FE
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
