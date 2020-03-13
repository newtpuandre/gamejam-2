package main

import (
	"image"
	"log"
	"os"

	_ "image/png"

	"github.com/hajimehoshi/ebiten"
)

const (
	screenWidth  = 1280
	screenHeight = 720
)

var (
	//Sprites and stuff
	player *ebiten.Image
	block  *ebiten.Image

	//Misc variables
	charX = 50
	charY = 300
	dx    = 3
	dy    = 10
)

func loadImageFile(filepath string) image.Image {
	imgFile, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}

	img, _, err := image.Decode(imgFile)
	if err != nil {
		log.Fatal(err)
	}

	return img
}

func loadImages() {

	player, _ = ebiten.NewImageFromImage(loadImageFile("./images/player.png"), ebiten.FilterDefault)
	block, _ = ebiten.NewImageFromImage(loadImageFile("./images/block1.png"), ebiten.FilterDefault)

}

func update(screen *ebiten.Image) error {

	if ebiten.IsDrawingSkipped() {
		return nil
	}

	if ebiten.IsKeyPressed(ebiten.KeyD) {
		charX += 3
	} else if ebiten.IsKeyPressed(ebiten.KeyA) {
		charX -= 3
	} else if ebiten.IsKeyPressed(ebiten.KeyW) {
		charY -= 3
	} /*else if ebiten.IsKeyPressed(ebiten.KeyS) {
		charY += 3
	} */

	// Add velocity
	//charX += 3
	charY += dy

	//Check if we are colliding with stuff
	//If we are dont apply vertical velocity

	//Player image options
	playerOptions := &ebiten.DrawImageOptions{}
	playerOptions.GeoM.Translate(float64(charX), float64(charY))

	//Block options (?)
	blockOptions := &ebiten.DrawImageOptions{}
	blockOptions.GeoM.Translate(float64(250), float64(250))

	//Draw the player
	screen.DrawImage(player, playerOptions)
	screen.DrawImage(block, blockOptions)

	return nil
}

func main() {
	loadImages()

	if err := ebiten.Run(update, screenWidth, screenHeight, 1, "Game"); err != nil {
		log.Fatal(err)
	}
}
