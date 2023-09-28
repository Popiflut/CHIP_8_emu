package emu

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"os"
)

type CPU struct {
	memory [4096]byte
	v      [16]uint8
	pc     uint16
	i      uint16
	dt     uint8
	st     uint8
	sp     uint16
	stack  [16]uint16
}

type Chip8 struct {
	cpu     CPU
	clavier Clavier
	screen  Screen
}

var (
	chip8  = Chip8{}
	ROM    []byte
	screen *ebiten.Image
)

func LoadROM(file []byte) {
	chip8.cpu.memory = [4096]byte{}
	n := 0x200
	for i := 0; i < len(file); i++ {
		chip8.cpu.memory[n] = file[i]
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
	//fmt.Println(chip8.clavier.Keyboard)
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

func StackPush(pc uint16) {
	chip8.cpu.stack[chip8.cpu.sp] = pc
	chip8.cpu.sp++
}

func StackPop() uint16 {
	chip8.cpu.sp--
	return chip8.cpu.stack[chip8.cpu.sp]
}
