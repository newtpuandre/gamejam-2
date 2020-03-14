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

	updateMovement(blocks)
	updateMovement(spikes)

	drawSprites(screen, blocks)
	drawSprites(screen, spikes)

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

func drawSprites(screen *ebiten.Image, s []Sprite) {

	for _, elem := range s {
		if elem.x > -64 && elem.x < screenWidth {
			spriteOptions := &ebiten.DrawImageOptions{}
			spriteOptions.GeoM.Translate(elem.x, elem.y)
			screen.DrawImage(elem.Image, spriteOptions)
		}
	}

}

func updateMovement(s []Sprite) {
	for i, elem := range s {
		if gameState > 0 {
			s[i].x -= elem.dx
		}
	}
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
