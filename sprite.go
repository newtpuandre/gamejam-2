package main

import "github.com/hajimehoshi/ebiten"

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
