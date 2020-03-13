package main

import (
	"log"
	"os"

	_ "image/png"

	"github.com/hajimehoshi/ebiten"
)

const (
	screenWidth  = 1280
	screenHeight = 720

	//Block and player size
	BLOCK_SIZE = 64
)

var (
	//Sprites and stuff
	player    Sprite
	player2   Sprite
	portal    Sprite
	startGame Sprite

	blocks []Sprite
	spikes []Sprite

	gameState int

	playerJumping bool
	ableToJump    bool
)

func update(screen *ebiten.Image) error {

	if ebiten.IsDrawingSkipped() {
		return nil
	}

	if gameState != 0 {
		//Handle keyboard input
		if ebiten.IsKeyPressed(ebiten.KeySpace) {

			if !playerJumping && ableToJump {
				playerJumping = true
				ableToJump = false
			}

		}

		if playerJumping == true {

			player.y -= player.dy

			if player.y < 285 {
				playerJumping = false
			}

		} else if player.y < 436 && playerJumping == false {
			player.y += player.dy - 1
			if player.y >= 436 {
				ableToJump = true
			}
		}

	}

	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		os.Exit(1)
	}

	// Blocks
	for i, elem := range blocks {
		if gameState > 0 {
			blocks[i].x -= elem.dx
		}
	}

	for _, elem := range blocks {
		spriteOptions := &ebiten.DrawImageOptions{}
		spriteOptions.GeoM.Translate(elem.x, elem.y)
		screen.DrawImage(elem.Image, spriteOptions)
	}

	//Spikes
	for i, elem := range spikes {
		if gameState > 0 {
			spikes[i].x -= elem.dx
		}
	}

	for _, elem := range spikes {
		spriteOptions := &ebiten.DrawImageOptions{}
		spriteOptions.GeoM.Translate(elem.x, elem.y)
		screen.DrawImage(elem.Image, spriteOptions)
	}

	for _, elem := range spikes {

		if doColide(player, elem) {
			log.Println("test")
		}

	}

	//Player image options
	playerOptions := &ebiten.DrawImageOptions{}
	playerOptions.GeoM.Translate(player.x, player.y)

	//Draw the player
	screen.DrawImage(player.Image, playerOptions)

	if gameState == 0 {
		if ebiten.IsKeyPressed(ebiten.KeyEnter) {
			gameState = 1
		}

		StartGameOptions := &ebiten.DrawImageOptions{}
		StartGameOptions.GeoM.Translate(startGame.x, startGame.y)
		screen.DrawImage(startGame.Image, StartGameOptions)

	}

	return nil
}

func main() {
	gameState = 0

	ableToJump = true

	loadImages()

	player.x = 250
	player.y = 436
	player.dx = 3
	player.dy = 8

	if err := ebiten.Run(update, screenWidth, screenHeight, 1, "Game jam: Shapeshifting game"); err != nil {
		log.Fatal(err)
	}
}
