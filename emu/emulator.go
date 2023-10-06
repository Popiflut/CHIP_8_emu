package emu

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"main/emu/Utils"
	"main/emu/VAR"
	"os"
	"strings"
)

// LoadROM -> charge le ROM dans la memoire
func LoadROM(file []byte) {
	VAR.CHIP8.Cpu.Memory = [4096]byte{}
	n := 0x200
	for i := 0; i < len(file); i++ {
		VAR.CHIP8.Cpu.Memory[n] = file[i]
		n++
	}
}

// Start -> lance l'emulator
func Start() error {
	file, err := os.ReadFile(os.Args[1])
	checkArgs()
	if err != nil {
		return err
	}
	VAR.ROM = file
	Init()
	fmt.Println("Init...")
	//fmt.Println(VAR.clavier.Keyboard)
	VAR.CHIP8.Screen.Mapscreen = [64][32]uint8{}
	LoadROM(file)
	VAR.CHIP8.Cpu.V = [16]uint8{}
	VAR.CHIP8.Cpu.Pc = 0x1FE
	fmt.Println("Loading ROM...")
	fmt.Println("ROM size: ", len(file))
	fmt.Println("Chip8 Emulator")
	fmt.Println("ROM: ", os.Args[1])
	fmt.Println("Start")
	LoadProgram()
	return nil
}

func checkArgs() {
	if len(os.Args) > 3 {
		fmt.Println("Usage: ./chip8 <ROM> [params]")
		os.Exit(0)
	} else if len(os.Args) == 3 {
		ebiten.SetTPS(60)
		for i := 2; i < len(os.Args); i++ {
			if strings.Contains(os.Args[i], "-TPS=") {
				ebiten.SetTPS(Utils.AtoI(os.Args[i][5:]))
			}
		}
	}
}
