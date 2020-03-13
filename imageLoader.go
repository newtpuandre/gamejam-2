package main

import (
	"image"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten"
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

	player.Image, _ = ebiten.NewImageFromImage(loadImageFile("./images/player.png"), ebiten.FilterDefault)
	player2.Image, _ = ebiten.NewImageFromImage(loadImageFile("./images/player_fly.png"), ebiten.FilterDefault)

	startGame.Image, _ = ebiten.NewImageFromImage(loadImageFile("./images/startgame.png"), ebiten.FilterDefault)
	startGame.x = 375
	startGame.y = 300

	var tempSprite Sprite

	tempSprite.Image, _ = ebiten.NewImageFromImage(loadImageFile("./images/block1.png"), ebiten.FilterDefault)

	//Add lots of blocks
	for i := 0; i < 50; i++ {
		tempSprite.x = 250 + float64((BLOCK_SIZE * i))
		tempSprite.y = 500
		tempSprite.dx = 5
		tempSprite.dy = 0
		blocks = append(blocks, tempSprite)
	}

	var tempSpriteSpike Sprite

	tempSpriteSpike.Image, _ = ebiten.NewImageFromImage(loadImageFile("./images/spikes.png"), ebiten.FilterDefault)
	for i := 0; i < 10; i++ {
		tempSpriteSpike.x = 442 + float64((BLOCK_SIZE*i)*4)

		tempSpriteSpike.y = 436
		tempSpriteSpike.dx = 5
		spikes = append(spikes, tempSpriteSpike)
	}

}
