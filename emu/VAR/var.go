package VAR

import (
	"github.com/hajimehoshi/ebiten/v2"
	"time"
)

// CHIP8 is the main struct of the emulator
// It contains the CPU, the keyboard and the screen
type Chip8 struct {
	Cpu     CPUs
	Clavier Claviers
	Screen  Screens
}

// Claviers is the keyboard struct
// It contains the state of the keys
type Claviers struct {
	IsPressed [16]bool
}

// Screens is the screen struct
// It contains the state of the screen
type Screens struct {
	Mapscreen [64][32]uint8
}

// CPUs is the CPU struct
// It contains the state of the CPU
// It contains the memory, the registers, the program counter, the index register, the delay timer, the sound timer, the stack pointer and the stack
type CPUs struct {
	Memory    [4096]byte
	V         [16]uint8
	Pc        uint16
	I         uint16
	Dt        uint8
	TimeStart time.Time
	St        uint8
	Sp        uint16
	Stack     [16]uint16
}

const (
	// ScreenWidth is the width of the screen
	ScreenWidth = 640
	// ScreenHeight is the height of the screen
	ScreenHeight = 320
	// ResolWidth is the width of the resolution
	ResolWidth = 64
	// ResolHeight is the height of the resolution
	ResolHeight = 32
)

var (
	// CHIP8 is the main struct of the emulator
	CHIP8 = Chip8{}
	// ROM is the file containing the game
	ROM []byte
	// screen is the image for ebiten
	screen *ebiten.Image
)

// GetKey returns the state of the key
func (kb *Claviers) GetKey(key byte) bool {
	return kb.IsPressed[key]
}
