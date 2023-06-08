package objects

import (
	"math/rand"
)

type Food struct {
	Point Point
}

func NewFood() *Food {
	return &Food{
		Point: Point{
			X: rand.Intn(SCREEN_WIDTH/10) * 10,
			Y: rand.Intn(SCREEN_HEIGHT/10) * 10,
		},
	}
}
