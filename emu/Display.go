package emu

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
	"log"
)

type Screen struct {
	Mapscreen [64][32]uint8
}

type Console struct {
	IN      string
	OUT     string
	command string
}

const (
	ScreenWidth  = 640
	screenHeight = 320
	ResolWidth   = 64
	ResolHeight  = 32
)

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
		CHIP8.Cpu.memory[i] = fontSet[i]
	}
	for i := 0; i < len(CHIP8.clavier.isPressed); i++ {
		CHIP8.clavier.isPressed[i] = false
	}
}

// NewConsole initialise un nouveau jeu.
func NewConsole() *Console {
	console := &Console{}
	return console
}

func (g *Console) Update() error {
	CHIP8.Cpu.Pc += 2
	fmt.Printf("cp = %02X:0x%04X: ", CHIP8.Cpu.Pc, (uint16(CHIP8.Cpu.memory[CHIP8.Cpu.Pc])<<8)|uint16(CHIP8.Cpu.memory[CHIP8.Cpu.Pc+1]))
	CHIP8.Cpu.Interpreter((uint16(CHIP8.Cpu.memory[CHIP8.Cpu.Pc]) << 8) | uint16(CHIP8.Cpu.memory[CHIP8.Cpu.Pc+1]))
	RefreshKeyBoard()
	return nil
}

func (g *Console) Draw(screen *ebiten.Image) {
	for x, row := range CHIP8.Screen.Mapscreen {
		for y, pixel := range row {
			if pixel == 1 {
				screen.Set(x, y, color.White)
			} else {
				screen.Set(x, y, color.Black)
			}

		}
	}
}

func (g *Console) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ResolWidth, ResolHeight
}

func LoadProgram() {
	//s = new
	ebiten.SetWindowSize(ScreenWidth, screenHeight)
	ebiten.SetWindowTitle("CHIP-8 Console")
	console := NewConsole()
	ebiten.SetTPS(60)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	if err := ebiten.RunGame(console); err != nil {
		log.Fatal(err)
	}
}
