package emu

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
	"log"
	"main/emu/VAR"
)

// Init initialise the font set and the keyboard
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
		VAR.CHIP8.Cpu.Memory[i] = fontSet[i]
	}
	for i := 0; i < len(VAR.CHIP8.Clavier.IsPressed); i++ {
		VAR.CHIP8.Clavier.IsPressed[i] = false
	}
}

// Consoles is the console struct
// It contains the input, the output and the command
type Consoles struct {
	IN      string
	OUT     string
	command string
}

// NewConsole initialise new emulator.
func NewConsole() *Consoles {
	console := &Consoles{}
	return console
}

// Update update the emulator.
func (g *Consoles) Update() error {
	VAR.CHIP8.Cpu.Pc += 2
	fmt.Printf("cp = %02X:0x%04X: ", VAR.CHIP8.Cpu.Pc, (uint16(VAR.CHIP8.Cpu.Memory[VAR.CHIP8.Cpu.Pc])<<8)|uint16(VAR.CHIP8.Cpu.Memory[VAR.CHIP8.Cpu.Pc+1]))
	VAR.CHIP8.Cpu.Interpreter((uint16(VAR.CHIP8.Cpu.Memory[VAR.CHIP8.Cpu.Pc]) << 8) | uint16(VAR.CHIP8.Cpu.Memory[VAR.CHIP8.Cpu.Pc+1]))
	RefreshKeyBoard()
	return nil
}

// Draw print the pixels on the screen.
func (g *Consoles) Draw(screen *ebiten.Image) {
	for x, row := range VAR.CHIP8.Screen.Mapscreen {
		for y, pixel := range row {
			if pixel == 1 {
				screen.Set(x, y, color.White)
			} else {
				screen.Set(x, y, color.Black)
			}
		}
	}
}

// Layout is the layout of the screen.
func (g *Consoles) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return VAR.ResolWidth, VAR.ResolHeight
}

// LoadProgram load the front program.
func LoadProgram() {
	//s = new
	ebiten.SetWindowSize(VAR.ScreenWidth, VAR.ScreenHeight)
	ebiten.SetWindowTitle("CHIP-8 Console")
	console := NewConsole()
	ebiten.SetTPS(60)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	if err := ebiten.RunGame(console); err != nil {
		log.Fatal(err)
	}
}
