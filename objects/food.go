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
			X: (rand.Intn(76-5) + 5) * 10,
			Y: (rand.Intn(56-17) + 17) * 10,
		},
	}
}
