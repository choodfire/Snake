package objects

import (
	"errors"
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

func (s *Snake) CheckBorders() error {
	head := s.Body[0]

	if head.X > 800 || head.X < 0 || head.Y > 600 || head.Y < 0 {
		return errors.New("snake hit border")
	}

	return nil
}
