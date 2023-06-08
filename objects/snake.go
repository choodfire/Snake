package objects

import (
	"fmt"
)

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
	newBody := make([]Point, 0, 10)
	newBody = append(newBody, Point{SCREEN_WIDTH / 4, SCREEN_HEIGHT / 2})

	return &Snake{
		Body:      newBody,
		Direction: Right,
		Length:    1,
	}
}

func (s *Snake) Move() {
	xMove := 0
	yMove := 0
	if s.Direction == Up {
		xMove = 0
		yMove = -10
	} else if s.Direction == Down {
		xMove = 0
		yMove = 10
	} else if s.Direction == Left {
		xMove = -10
		yMove = 0
	} else if s.Direction == Right {
		xMove = 10
		yMove = 0
	}

	for i := 0; i < len(s.Body); i++ {
		s.Body[i].X += xMove
		s.Body[i].Y += yMove
	}
}

func (s *Snake) ConsumeFood() {
	fmt.Println("i AtE fOoD")
}
