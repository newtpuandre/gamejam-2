package main

import (
	"fmt"
	"log"
	"os"

	"image/color"
	_ "image/png"

	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/text"
	"golang.org/x/image/font"
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
	startGame Sprite

	blocks    []Sprite
	flyBlocks []Sprite
	spikes    []Sprite
	portals   []Sprite

	gameState int

	arcadeFont font.Face

	playerJumping  bool
	ableToJump     bool
	score          float64
	isFlying       bool
	isFlyingJump   bool
	blockPortalCol bool

	counter int
	points  int64
)

func update(screen *ebiten.Image) error {
	if ebiten.IsDrawingSkipped() {
		return nil
	}
	screen.Fill(color.RGBA{0x80, 0xa0, 0xc0, 0xff})

	if isFlying {
		isFlyingJump = false
	}

	if gameState != 0 {
		points++
		if blockPortalCol {
			counter++
		}
		//Handle keyboard input
		if ebiten.IsKeyPressed(ebiten.KeySpace) {

			if (!playerJumping && ableToJump) && !isFlying {
				playerJumping = true
				ableToJump = false
			}

			if isFlying {
				isFlyingJump = true
			}

		}

		if !isFlying {
			if playerJumping {

				player.y -= player.dy

				if player.y < 244 {
					playerJumping = false
				}

			} else if player.y < 436 && playerJumping == false {
				player.y += player.dy - 1
				if player.y >= 436 {
					ableToJump = true
				}
			}

		} else {
			if !isFlyingJump {
				player.y += player.dy
			} else {
				player.y -= player.dy
			}
		}

	}

	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		os.Exit(1)
	}

	updateMovement(blocks)
	updateMovement(spikes)
	updateMovement(portals)
	updateMovement(flyBlocks)

	blockMove(blocks)
	blockMove(flyBlocks)

	if !isFlying {

		drawSprites(screen, blocks)
		drawSprites(screen, spikes)
	} else {

		drawSprites(screen, flyBlocks)
	}
	drawSprites(screen, portals)

	for _, elem := range spikes {

		if doColide(player, elem) {
			log.Println("We are colliding", player.x, elem.x)
		}

	}

	if !blockPortalCol {
		for _, elem := range portals {
			if doColide(player, elem) {
				log.Println("ShapeSHIFT")
				isFlying = !isFlying
				blockPortalCol = true
			}
		}
	}

	//Player image options
	playerOptions := &ebiten.DrawImageOptions{}
	playerOptions.GeoM.Translate(player.x, player.y)

	//Draw the player
	if isFlying {
		screen.DrawImage(player.SecondaryImage, playerOptions)

	} else {
		screen.DrawImage(player.Image, playerOptions)
	}

	if gameState == 0 {
		if ebiten.IsKeyPressed(ebiten.KeyEnter) {
			gameState = 1
		}

		StartGameOptions := &ebiten.DrawImageOptions{}
		StartGameOptions.GeoM.Translate(startGame.x, startGame.y)
		screen.DrawImage(startGame.Image, StartGameOptions)

	}

	//Reset blockPortalCol
	if blockPortalCol && counter > 100 {
		blockPortalCol = false
		counter = 0
	}

	scoreStr := fmt.Sprintf("%04d", points)
	text.Draw(screen, scoreStr, arcadeFont, screenWidth-len(scoreStr)*32, 32, color.White)

	return nil
}

func main() {
	tt, err := truetype.Parse(fonts.ArcadeN_ttf)
	if err != nil {
		log.Fatal(err)
	}
	arcadeFont = truetype.NewFace(tt, &truetype.Options{
		Size:    32,
		DPI:     72,
		Hinting: font.HintingFull,
	})
	gameState = 0

	ableToJump = true
	blockPortalCol = false

	loadImages()

	player.x = 250
	player.y = 436
	player.dx = 3
	player.dy = 8

	if err := ebiten.Run(update, screenWidth, screenHeight, 1, "Game jam: Shapeshifting game"); err != nil {
		log.Fatal(err)
	}
}
