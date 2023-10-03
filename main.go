package main

import (
	"fmt"
	"main/emu"
)

func main() {
	err := emu.Start()
	if err != nil {
		fmt.Println("ERROR system start")
		return
	}
}
