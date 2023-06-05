package objects

import "math/rand"

type Food struct {
	point Point
}

func NewFood() *Food {
	return &Food{
		point: Point{
			X: rand.Intn(800),
			Y: rand.Intn(600),
		},
	}
}
