package emu

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"main/emu/Utils"
	"main/emu/VAR"
	"os"
	"strings"
	"time"
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
	fmt.Println("-ROM: ", "-Name: ", os.Args[1], "-Size: ", len(file))
	VAR.Context = audio.NewContext(44100)
	LoadProgram()
	return nil
}

func checkArgs() {
	ebiten.SetTPS(60)
	VAR.CHIP8.Screen.TPS = 8
	if len(os.Args) > 4 {
		fmt.Println("Usage: ./chip8 <ROM> [params]")
		os.Exit(0)
	} else if len(os.Args) > 2 {
		print("-Params: ")
		for i := 2; i < len(os.Args); i++ {
			if strings.Contains(os.Args[i], "-CPU_HZ=") {
				ebiten.SetTPS(Utils.AtoI(os.Args[i][8:]))

				fmt.Print(" --Set_CPU_Hz: ", Utils.AtoI(os.Args[i][8:]), "Hz")
			} else if strings.Contains(os.Args[i], "-TPS=") {
				VAR.CHIP8.Screen.TPS = time.Duration(Utils.AtoI(os.Args[i][5:]))

				fmt.Print(" --Set_TPS: ", Utils.AtoI(os.Args[i][5:]), "fps")
			}
		}
		fmt.Println()
	}
}
