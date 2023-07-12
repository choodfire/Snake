package objects

import (
	"math/rand"
)

type Food struct {
	Point Point
}

func NewFood() *Food {
	min := (LEFT_BORDER + SQUARE_SIZE) / SQUARE_SIZE
	max := (UPPER_BORDER + SQUARE_SIZE) / SQUARE_SIZE

	return &Food{
		Point: Point{
			X: (rand.Intn((RIGHT_BORDER/SQUARE_SIZE)-min) + min) * SQUARE_SIZE,
			Y: (rand.Intn((BOTTOM_BORDER/SQUARE_SIZE)-max) + max) * SQUARE_SIZE,
		},
	}
}
