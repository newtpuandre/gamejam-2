package main

import "github.com/hajimehoshi/ebiten"

type Sprite struct {
	x                    float64
	y                    float64
	dx                   float64
	dy                   float64
	Image                *ebiten.Image
	ImageInvuln          *ebiten.Image
	SecondaryImage       *ebiten.Image
	SecondaryImageInvuln *ebiten.Image
}

func doColide(s1 Sprite, s2 Sprite) bool {

	//We have a point in space representing the start of the rectangle.
	// x,y is this point. X + BLOCK_SIZE is top right of the rectangle
	// Y + BLOCK_SIZE is bottom left.
	// Bottom right is Y + ( top left - top right)

	/*s1lx := s1.x                 //Top left
	s1ly := s1.y + BLOCK_SIZE    //Bottom left
	s1rx := s1.x + BLOCK_SIZE    // Top Right
	s1ry := s1.y + (s1lx - s1ly) // Bottom right

	s2lx := s2.x                 //Top left
	s2ly := s2.y + BLOCK_SIZE    //Bottom left
	s2rx := s2.x + BLOCK_SIZE    // Top Right
	s2ry := s2.y + (s2lx - s2ly) // Bottom right

	if s1lx > s2rx || s2lx > s1rx {
		return false
	}

	if s1ly < s2ry || s2ly < s1ry {
		return false
	}

	return true*/

	if (s1.x-s2.x < 64 && s2.x-s1.x < 64) && (s1.y-s2.y < 64 && s2.y-s1.y < 64) {
		return true
	}

	return false

}

// Moves the blocks that are out of screen behind the user to the front.
func blockMove(s []Sprite) {
	for i, elem := range s {
		if elem.x < -BLOCK_SIZE {
			s[i].x = screenWidth + BLOCK_SIZE
		}
	}
}

func updateMovement(s []Sprite) {
	for i := range s {
		if gameState > 0 {
			s[i].x -= globaldx
		}
	}
}

func drawSprites(screen *ebiten.Image, s []Sprite) {

	for _, elem := range s {
		if elem.x > -BLOCK_SIZE && elem.x < screenWidth {
			spriteOptions := &ebiten.DrawImageOptions{}
			spriteOptions.GeoM.Translate(elem.x, elem.y)
			screen.DrawImage(elem.Image, spriteOptions)
		}
	}

}

func drawPortalSprites(screen *ebiten.Image, s []Sprite) {

	for _, elem := range s {
		if elem.x > -BLOCK_SIZE && elem.x < screenWidth {
			spriteOptions := &ebiten.DrawImageOptions{}
			spriteOptions.GeoM.Translate(elem.x, elem.y)
			if isFlying {
				screen.DrawImage(elem.SecondaryImage, spriteOptions)
			} else {
				screen.DrawImage(elem.Image, spriteOptions)
			}
		}
	}
}
