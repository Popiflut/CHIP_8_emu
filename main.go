package main

//faire une console ds une nouvelle fenetre avec ebiten

import "github.com/hajimehoshi/ebiten"

func main() {
	//afficher une fenetre
	err := ebiten.Run(update, 320, 240, 2, "Hello world")
	if err != nil {
		return
	}
}

func update(image *ebiten.Image) error {
	return nil
}
