package emu

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"os"
)

type CPU struct {
	memory [4096]byte
	v      [16]uint8
	Pc     uint16
	i      uint16
	dt     uint8
	st     uint8
	sp     uint16
	stack  [16]uint16
}

type Chip8 struct {
	Cpu     CPU
	clavier Clavier
	Screen  Screen
}

var (
	CHIP8  = Chip8{}
	ROM    []byte
	screen *ebiten.Image
)

func LoadROM(file []byte) {
	CHIP8.Cpu.memory = [4096]byte{}
	n := 0x200
	for i := 0; i < len(file); i++ {
		CHIP8.Cpu.memory[n] = file[i]
		n++
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
	//fmt.Println(CHIP8.clavier.Keyboard)
	CHIP8.Screen.Mapscreen = [64][32]uint8{}
	LoadROM(file)
	CHIP8.Cpu.Pc = 0x1FE
	ebiten.SetTPS(60)
	fmt.Println("Loading ROM...")
	fmt.Println("ROM size: ", len(file))
	fmt.Println("Chip8 Emulator")
	fmt.Println("ROM: ", os.Args[1])
	fmt.Println("Start")
	LoadProgram()
	return nil
}

func StackPush(pc uint16) {
	CHIP8.Cpu.stack[CHIP8.Cpu.sp] = pc
	CHIP8.Cpu.sp++
}

func StackPop() uint16 {
	CHIP8.Cpu.sp--
	return CHIP8.Cpu.stack[CHIP8.Cpu.sp]
}
