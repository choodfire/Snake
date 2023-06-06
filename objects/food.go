package objects

import "math/rand"

type Food struct {
	Point Point
}

func NewFood() *Food {
	return &Food{
		Point: Point{
			X: rand.Intn(800),
			Y: rand.Intn(600),
		},
	}
}
