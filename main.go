package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
	"log"
	"os"
)

var (
	chip8 = Chip8{}
	ROM   = []byte{}
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

// NewConsole initialise un nouveau jeu.
func NewConsole() *Console {
	console := &Console{}
	return console
}

// Update met à jour l'état du jeu à chaque trame.
func (g *Console) Update() error {
	chip8.cpu.pc += 2
	chip8.cpu.Interpreter((uint16(chip8.cpu.memory[chip8.cpu.pc]) << 8) | uint16(chip8.cpu.memory[chip8.cpu.pc+1]))
	if chip8.cpu.pc >= len(ROM)+0x200 {
		os.Exit(0)
	}
	return nil
}

// Draw dessine le jeu sur l'écran
func (g *Console) Draw(screen *ebiten.Image) {
	for x := 0; x < len(chip8.screen.s); x++ {
		for y := 0; y < len(chip8.screen.s[x]); y++ {
			//ebitenutil.DrawRect(screen, float64(x*(screenWidth/64)), float64(y*(screenHeight/32)), 1, 1, chip8.screen.s[x][y])
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

//func Emu() {
//	for i := 0x200; i < len(chip8.cpu.memory)-1; i += 2 {
//		//fmt.Printf("0x%04X: ", (uint16(chip8.cpu.memory[i])<<8)|uint16(chip8.cpu.memory[i+1]))
//		jmp := Interpreter((uint16(chip8.cpu.memory[i]) << 8) | uint16(chip8.cpu.memory[i+1]))
//		if jmp != 0 {
//			fmt.Println(int(jmp & 0x0FFF))
//			fmt.Printf("JMP i =%d and now i = ", i)
//			i = int(jmp & 0x0FFF)
//			fmt.Println(i)
//		}
//	}
//}

func (cpu CPU) Interpreter(b uint16) uint16 {
	switch b & 0xF000 {
	case 0x0000:
		switch b & 0x000F {
		case 0x0000:
			fmt.Println("CLS")
			//remplacer toutes les valeurs de screen[][] pour du noir
		case 0x000E:
			fmt.Println("RET")
		}
	case 0x1000:
		chip8.cpu.pc = int(b & 0x0FFF)
	case 0x2000:
		fmt.Println("CALL addr")
	case 0x3000:
		fmt.Println("SE Vx, byte")
	case 0x4000:
		fmt.Println("SNE Vx, byte")
	case 0x5000:
		fmt.Println(" SE Vx, Vy")
	case 0x6000:
		fmt.Println("LD Vx, byte")
	case 0x7000:
		fmt.Println("ADD Vx, byte")
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
	//LoadProgram()
	file, err := os.ReadFile(os.Args[1])
	if err != nil {
		return err
	}
	ROM = file
	LoadROM(file)
	chip8.cpu.pc = 0x198
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
