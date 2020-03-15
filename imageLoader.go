package main

import (
	"image"
	"log"
	"math/rand"
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

	player.Image, _ = ebiten.NewImageFromImage(loadImageFile("./images/player_2.png"), ebiten.FilterDefault)
	player.ImageInvuln, _ = ebiten.NewImageFromImage(loadImageFile("./images/player_2_invuln.png"), ebiten.FilterDefault)
	player.SecondaryImage, _ = ebiten.NewImageFromImage(loadImageFile("./images/player_2_fly.png"), ebiten.FilterDefault)
	player.SecondaryImageInvuln, _ = ebiten.NewImageFromImage(loadImageFile("./images/player_2_fly_invuln.png"), ebiten.FilterDefault)
	player.ImageDeadFly, _ = ebiten.NewImageFromImage(loadImageFile("./images/player_2_fly_dead.png"), ebiten.FilterDefault)
	player.ImageDead, _ = ebiten.NewImageFromImage(loadImageFile("./images/player_2_dead.png"), ebiten.FilterDefault)

	startGame.Image, _ = ebiten.NewImageFromImage(loadImageFile("./images/startgame.png"), ebiten.FilterDefault)
	startGame.x = 375
	startGame.y = 300

	var tempSprite Sprite

	tempSprite.Image, _ = ebiten.NewImageFromImage(loadImageFile("./images/block1.png"), ebiten.FilterDefault)
	tempSprite.ImageVariation2, _ = ebiten.NewImageFromImage(loadImageFile("./images/block2.png"), ebiten.FilterDefault)
	tempSprite.ImageVariation3, _ = ebiten.NewImageFromImage(loadImageFile("./images/block3.png"), ebiten.FilterDefault)
	tempSprite.ImageVariation4, _ = ebiten.NewImageFromImage(loadImageFile("./images/block4.png"), ebiten.FilterDefault)

	//Add lots of blocks
	for i := 0; i < 22; i++ {
		tempSprite.x = 250 + float64((BLOCK_SIZE * i))
		tempSprite.y = 500
		tempSprite.dx = 5
		tempSprite.dy = 0
		tempSprite.draw = true
		blocks = append(blocks, tempSprite)
	}

	//Add lots of blocks
	for i := 0; i < 22; i++ {
		tempSprite.x = 1658 + float64((BLOCK_SIZE*i)*3)
		tempSprite.y = 200 * float64(rand.Intn(4))
		tempSprite.dx = 5
		tempSprite.dy = 0
		tempSprite.draw = true
		flyBlocks = append(flyBlocks, tempSprite)
	}

	var tempSpriteSpike Sprite

	tempSpriteSpike.Image, _ = ebiten.NewImageFromImage(loadImageFile("./images/spikes.png"), ebiten.FilterDefault)
	for i := 0; i < 100; i++ {
		tempSpriteSpike.x = 442 + float64((BLOCK_SIZE*i)*6)

		tempSpriteSpike.y = 436
		tempSpriteSpike.dx = 5
		tempSpriteSpike.draw = true
		spikes = append(spikes, tempSpriteSpike)
	}

	var tempSpritePortal Sprite
	tempSpritePortal.Image, _ = ebiten.NewImageFromImage(loadImageFile("./images/portal.png"), ebiten.FilterDefault)
	tempSpritePortal.SecondaryImage, _ = ebiten.NewImageFromImage(loadImageFile("./images/portal_2.png"), ebiten.FilterDefault)
	for i := 0; i < 20; i++ {
		tempSpritePortal.x = float64(1402 * (4 + i))
		tempSpritePortal.y = 308
		tempSpritePortal.dx = 5
		tempSpritePortal.draw = true
		portals = append(portals, tempSpritePortal)
	}

	var tempSpriteStar Sprite
	tempSpriteStar.Image, _ = ebiten.NewImageFromImage(loadImageFile("./images/star.png"), ebiten.FilterDefault)
	for i := 0; i < 20; i++ {
		tempSpriteStar.x = 384 + 1280*float64((2*i))
		tempSpriteStar.y = 436
		tempSpriteStar.dx = 5
		tempSpriteStar.draw = true
		stars = append(stars, tempSpriteStar)
	}

	for i := 0; i < 22; i++ {
		tempSpriteStar.x = 384 + 1280*float64((2*i))
		tempSpriteStar.y = 200 * float64(rand.Intn(4))
		tempSpriteStar.dx = 5
		tempSpriteStar.draw = true
		flyStars = append(flyStars, tempSpriteStar)
	}

}
