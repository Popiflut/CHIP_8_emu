package emu

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
	"log"
)

type Screen struct {
	mapscreen [64][32]uint8
}

type Console struct {
	IN      string
	OUT     string
	command string
}

const (
	screenWidth  = 640
	screenHeight = 320
	resolWidth   = 64
	resolHeight  = 32
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
		chip8.cpu.memory[i] = fontSet[i]
	}
	for i := 0; i < len(chip8.clavier.isPressed); i++ {
		chip8.clavier.isPressed[i] = false
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
