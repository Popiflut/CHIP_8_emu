package main

import (
	"fmt"
	_ "github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"main/emu"
)

func main() {
	err := emu.Start()
	if err != nil {
		fmt.Println("ERROR system start")
		return
	}
}
