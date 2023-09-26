package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 640
	screenHeight = 480
	resolWidth   = 400
	resolHeight  = 364
)

// Game représente l'état du jeu.
type Game struct {
	menu int
}

// NewGame initialise un nouveau jeu.
func NewGame() *Game {
	game := &Game{}
	return game
}

// Update met à jour l'état du jeu à chaque trame.
func (g *Game) Update() error {
	// Mettez à jour ici la logique du jeu, par exemple, la position des objets, la détection de collisions, etc
	return nil
}

// Draw dessine le jeu sur l'écran.
func (g *Game) Draw(screen *ebiten.Image) {
	// Dessinez ici les objets de jeu sur l'écran en utilisant les fonctions d'Ebiten.
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return resolWidth, resolHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Mon Jeu Ebiten")

	game := NewGame()
	ebiten.SetTPS(60)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
