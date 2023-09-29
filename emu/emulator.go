package emu

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"main/emu/VAR"
	"os"
)

// LoadROM -> Load the ROM in memory
func LoadROM(file []byte) {
	VAR.CHIP8.Cpu.Memory = [4096]byte{}
	n := 0x200
	for i := 0; i < len(file); i++ {
		VAR.CHIP8.Cpu.Memory[n] = file[i]
		n++
	}
}

// Start -> Start the emulator
func Start() error {
	file, err := os.ReadFile(os.Args[1])
	if err != nil {
		return err
	}
	VAR.ROM = file
	Init()
	fmt.Println("Init...")
	//fmt.Println(VAR.clavier.Keyboard)
	VAR.CHIP8.Screen.Mapscreen = [64][32]uint8{}
	LoadROM(file)
	VAR.CHIP8.Cpu.Pc = 0x1FE
	ebiten.SetTPS(60)
	fmt.Println("Loading ROM...")
	fmt.Println("ROM size: ", len(file))
	fmt.Println("Chip8 Emulator")
	fmt.Println("ROM: ", os.Args[1])
	fmt.Println("Start")
	LoadProgram()
	return nil
}
