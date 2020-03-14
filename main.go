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

	globaldx float64

	gameState int

	arcadeFont font.Face

	playerJumping bool
	ableToJump    bool
	score         float64
	isFlying      bool
	isFlyingJump  bool

	blockPortalCol bool

	counter int

	points    int64
	highscore int64

	invulnerability bool
	invulnCounter   int
)

func update(screen *ebiten.Image) error {
	if ebiten.IsDrawingSkipped() {
		return nil
	}
	screen.Fill(color.RGBA{0x80, 0xa0, 0xc0, 0xff})

	if isFlying {
		isFlyingJump = false
	}

	if gameState == 1 {
		points++
		if blockPortalCol {
			counter++
		}
		if invulnerability {
			invulnCounter++
		}

		globaldx += 0.001
		log.Println(globaldx)
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
				if player.y+BLOCK_SIZE < screenHeight {
					player.y += player.dy
				}
			} else {
				if player.y > 0 {
					player.y -= player.dy
				}
			}
		}

	}

	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		os.Exit(1)
	}

	if gameState == 1 {

		updateMovement(portals)

	}

	if !isFlying {
		if gameState == 1 {
			updateMovement(blocks)
			updateMovement(spikes)
			blockMove(blocks)
		}
		drawSprites(screen, blocks)
		drawSprites(screen, spikes)
		if !invulnerability {
			for _, elem := range spikes {

				if doColide(player, elem) {
					log.Println("We are colliding", player.x, elem.x)
					gameState = 2
				}

			}
		}
	} else {
		if gameState == 1 {
			updateMovement(flyBlocks)

			blockMove(flyBlocks)
		}
		drawSprites(screen, flyBlocks)
		if !invulnerability {

			for _, elem := range flyBlocks {

				if doColide(player, elem) {
					log.Println("We are colliding", player.x, elem.x)
					gameState = 2
				}

			}
		}
	}
	drawPortalSprites(screen, portals)

	if !blockPortalCol {
		for _, elem := range portals {
			if doColide(player, elem) {
				log.Println("ShapeSHIFT")
				isFlying = !isFlying
				blockPortalCol = true
				invulnerability = true
				points += 100
			}
		}
	}

	//Player image options
	playerOptions := &ebiten.DrawImageOptions{}
	playerOptions.GeoM.Translate(player.x, player.y)

	//Draw the player
	if isFlying {
		if invulnerability {
			screen.DrawImage(player.SecondaryImageInvuln, playerOptions)
		} else {

			screen.DrawImage(player.SecondaryImage, playerOptions)
		}

	} else {
		if invulnerability {
			screen.DrawImage(player.ImageInvuln, playerOptions)
		} else {
			screen.DrawImage(player.Image, playerOptions)
		}
	}

	if gameState == 0 {
		if ebiten.IsKeyPressed(ebiten.KeyEnter) {
			gameState = 1
		}

		StartGameOptions := &ebiten.DrawImageOptions{}
		StartGameOptions.GeoM.Translate(startGame.x, startGame.y)
		screen.DrawImage(startGame.Image, StartGameOptions)

	}

	if gameState == 2 {
		scoreStr := fmt.Sprintf("You died. Total Points: %07d", points)
		text.Draw(screen, scoreStr, arcadeFont, 145, 200, color.White)
		text.Draw(screen, "Press 'ENTER' to play again.", arcadeFont, 155, 250, color.White)
		if points > highscoreStruct.Highscore {

			highscore = points
		}
		writeHighscore()
		if ebiten.IsKeyPressed(ebiten.KeyEnter) {
			blocks = make([]Sprite, 0)
			spikes = make([]Sprite, 0)
			flyBlocks = make([]Sprite, 0)
			portals = make([]Sprite, 0)
			loadImages()
			player.x = 250
			player.y = 436
			player.dx = 3
			player.dy = 8
			gameState = 0
			ableToJump = true
			blockPortalCol = false
			isFlying = false
			isFlyingJump = false
			playerJumping = false
			invulnerability = false
			counter = 0
			points = 0
			globaldx = 5
		}
	}

	//Reset blockPortalCol
	if blockPortalCol && counter > 100 {
		blockPortalCol = false
		counter = 0
	}

	if invulnerability && invulnCounter > 100 {
		invulnerability = false
		invulnCounter = 0
	}

	scoreStr := fmt.Sprintf("%07d", points)
	highscoreStr := fmt.Sprintf("Highscore: %07d", highscore)
	text.Draw(screen, scoreStr, arcadeFont, screenWidth-len(scoreStr)*32, 32, color.White)
	text.Draw(screen, highscoreStr, arcadeFont, 0, 32, color.White)

	return nil
}

func main() {
	ConfigInit()
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
	invulnerability = false
	//isFlying = true

	loadImages()

	player.x = 250
	player.y = 436
	player.dx = 3
	player.dy = 8

	globaldx = 5

	if err := ebiten.Run(update, screenWidth, screenHeight, 1, "Game jam: Shapeshifting game"); err != nil {
		log.Fatal(err)
	}
}
