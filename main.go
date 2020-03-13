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

	//Block and player size
	BLOCK_SIZE = 64
)

var (
	//Sprites and stuff
	player    Sprite
	startGame Sprite

	sprites []Sprite

	gameState int

	playerJumping bool
	ableToJump    bool
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

	startGame.Image, _ = ebiten.NewImageFromImage(loadImageFile("./images/startgame.png"), ebiten.FilterDefault)
	startGame.x = 375
	startGame.y = 300

	var tempSprite Sprite

	tempSprite.Image, _ = ebiten.NewImageFromImage(loadImageFile("./images/block1.png"), ebiten.FilterDefault)

	//Add 5 blocks
	for i := 0; i < 100; i++ {
		tempSprite.x = 250 + float64((BLOCK_SIZE * i))
		tempSprite.y = 250
		tempSprite.dx = 5
		tempSprite.dy = 0
		sprites = append(sprites, tempSprite)
	}

}

func update(screen *ebiten.Image) error {

	if ebiten.IsDrawingSkipped() {
		return nil
	}

	if gameState != 0 {
		//Handle keyboard input
		if ebiten.IsKeyPressed(ebiten.KeySpace) {
			log.Println("Jumping...")
			if !playerJumping && ableToJump {
				playerJumping = true
				ableToJump = false
			}

		}

		if playerJumping == true {

			player.y -= player.dy

			if player.y < 58 {
				playerJumping = false
			}

		} else if player.y < 186 && playerJumping == false {
			player.y += player.dy + 2
			if player.y >= 186 {
				ableToJump = true
			}
		}

	}

	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		os.Exit(1)
	}

	log.Println(player.y)

	//We should only allow jumps and a different key to shapeshift

	colliding := false

	//Check for collisions
	// TODO: Collisions. Allows for holes in the ground etc.
	/*for _, elem := range sprites {

		if doColide(player, elem) {
			colliding = true
		}

	}*/

	if colliding {
		player.y += player.dy
	}

	//Check if we are colliding with stuff
	//If we are dont apply vertical velocity

	// Block updating
	for i, elem := range sprites {
		if gameState > 0 {
			sprites[i].x -= elem.dx
		}

	}

	//Block drawing
	for _, elem := range sprites {
		spriteOptions := &ebiten.DrawImageOptions{}
		spriteOptions.GeoM.Translate(elem.x, elem.y)
		screen.DrawImage(elem.Image, spriteOptions)
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
	gameState = 0 //Press enter to start

	ableToJump = true

	loadImages()

	player.x = 250
	player.y = 186
	player.dx = 3
	player.dy = 7

	if err := ebiten.Run(update, screenWidth, screenHeight, 1, "Game"); err != nil {
		log.Fatal(err)
	}
}

type Sprite struct {
	x     float64
	y     float64
	dx    float64
	dy    float64
	Image *ebiten.Image
}

func doColide(s1 Sprite, s2 Sprite) bool {

	//We have a point in space representing the start of the rectangle.
	// x,y is this point. X + BLOCK_SIZE is top right of the rectangle
	// Y + BLOCK_SIZE is bottom left.
	// Bottom right is Y + ( top left - top right)

	s1lx := s1.x                 //Top left
	s1ly := s1.y + BLOCK_SIZE    //Bottom left
	s1rx := s1.x + BLOCK_SIZE    // Top Right
	s1ry := s1.y + (s1lx - s1ly) // Bottom right

	s2lx := s2.x                 //Top left
	s2ly := s2.y                 //Bottom left
	s2rx := s2.x + BLOCK_SIZE    // Top Right
	s2ry := s2.y + (s2lx - s2ly) // Bottom right

	if s1lx > s2rx || s2lx > s1rx {
		return false
	}

	if s1ly < s2ry || s2ly < s1ry {
		return false
	}

	return true

}
