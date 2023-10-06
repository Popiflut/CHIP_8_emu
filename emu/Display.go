package emu

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
	"log"
	"main/emu/VAR"
	"time"
)

// Init initialise le font set et le keyboard
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

var (
	Draw bool
)

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

// Update update l'emulator.
func (g *Consoles) Update() error {
	if time.Now().Sub(VAR.CHIP8.Cpu.TimeStart) > time.Second/VAR.CHIP8.Screen.TPS { // when one second has past
		//if time.Now().Sub(VAR.CHIP8.Cpu.TimeStart) > time.Second/8 {
		if VAR.CHIP8.Cpu.Dt > 0 {
			VAR.CHIP8.Cpu.Dt -= 1
		}
		VAR.CHIP8.Cpu.TimeStart = time.Now()
		Draw = true
		if VAR.CHIP8.Cpu.SoundTimer > 0 {
			VAR.CHIP8.Cpu.SoundTimer -= 1
		}
	}
	VAR.CHIP8.Cpu.Pc += 2
	VAR.CHIP8.Cpu.Interpreter((uint16(VAR.CHIP8.Cpu.Memory[VAR.CHIP8.Cpu.Pc]) << 8) | uint16(VAR.CHIP8.Cpu.Memory[VAR.CHIP8.Cpu.Pc+1]))
	RefreshKeyBoard()

	// update audio
	var volume float64
	if VAR.CHIP8.Cpu.SoundTimer > 0 {
		volume = 1
	}
	VAR.CHIP8.Cpu.AudioPlayer.SetVolume(volume)
	return nil
}

// Draw print les pixels sur l'ecran.
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

// Layout est le layout de l'ecran.
func (g *Consoles) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return VAR.ResolWidth, VAR.ResolHeight
}

// LoadProgram load the front program.
func LoadProgram() {
	//s = new
	ebiten.SetWindowSize(VAR.ScreenWidth, VAR.ScreenHeight)
	ebiten.SetWindowTitle("CHIP-8 Console")
	console := NewConsole()
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	if err := ebiten.RunGame(console); err != nil {
		log.Fatal(err)
	}
}
