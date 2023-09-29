package VAR

import "github.com/hajimehoshi/ebiten/v2"

type Chip8 struct {
	Cpu     CPUs
	Clavier Claviers
	Screen  Screens
}

type Claviers struct {
	IsPressed [16]bool
}

type Screens struct {
	Mapscreen [64][32]uint8
}

type CPUs struct {
	Memory [4096]byte
	V      [16]uint8
	Pc     uint16
	I      uint16
	Dt     uint8
	St     uint8
	Sp     uint16
	Stack  [16]uint16
}

const (
	ScreenWidth  = 640
	ScreenHeight = 320
	ResolWidth   = 64
	ResolHeight  = 32
)

var (
	CHIP8  = Chip8{}
	ROM    []byte
	screen *ebiten.Image
)

func (kb *Claviers) GetKey(key byte) bool {
	return kb.IsPressed[key]
}
