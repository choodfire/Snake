package objects

import "fmt"

type Direction int

const (
	Left Direction = iota
	Right
	Up
	Down
)

type Snake struct {
	Body      []Point
	Direction Direction
	Length    int
}

func NewSnake() *Snake {
	return &Snake{
		Body:      []Point{{100, 300}},
		Direction: Right,
		Length:    1,
	}
}

func (s *Snake) Move() {
	xMove := 0
	yMove := 0
	if s.Direction == Up {
		xMove = 0
		yMove = -1
	} else if s.Direction == Down {
		xMove = 0
		yMove = 1
	} else if s.Direction == Left {
		xMove = -1
		yMove = 0
	} else if s.Direction == Right {
		xMove = 1
		yMove = 0
	}

	for i := 0; i < len(s.Body); i++ {
		s.Body[i].X += xMove
		s.Body[i].Y += yMove
	}

	fmt.Println(s.Body)
}
