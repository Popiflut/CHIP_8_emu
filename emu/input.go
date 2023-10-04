package emu

import (
	"github.com/hajimehoshi/ebiten/v2"
	"main/emu/VAR"
)

// RefreshKeyBoard -> Refresh the keyboard
func RefreshKeyBoard() {
	VAR.CHIP8.Clavier.IsPressed = [16]bool{
		ebiten.IsKeyPressed(ebiten.KeyX), //0
		ebiten.IsKeyPressed(ebiten.Key1), //1
		ebiten.IsKeyPressed(ebiten.Key2), //2
		ebiten.IsKeyPressed(ebiten.Key3), //3
		ebiten.IsKeyPressed(ebiten.KeyQ), //4
		ebiten.IsKeyPressed(ebiten.KeyW), //5
		ebiten.IsKeyPressed(ebiten.KeyE), //6
		ebiten.IsKeyPressed(ebiten.KeyA), //7
		ebiten.IsKeyPressed(ebiten.KeyS), //8
		ebiten.IsKeyPressed(ebiten.KeyD), //9
		ebiten.IsKeyPressed(ebiten.KeyZ), //A
		ebiten.IsKeyPressed(ebiten.KeyC), //B
		ebiten.IsKeyPressed(ebiten.Key4), //C
		ebiten.IsKeyPressed(ebiten.KeyR), //D
		ebiten.IsKeyPressed(ebiten.KeyF), //E
		ebiten.IsKeyPressed(ebiten.KeyV), //F
	}
}
