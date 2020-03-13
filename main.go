package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"
)

const (
	screenWidth  = 200
	screenHeight = 200
)

func update(screen *ebiten.Image) error {
	if ebiten.IsDrawingSkipped() {
		return nil
	}

	return nil
}

func main() {
	if err := ebiten.Run(update, screenWidth, screenHeight, 2, "Game"); err != nil {
		log.Fatal(err)
	}
}
