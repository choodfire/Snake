package objects

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
)

type Direction int

const (
	Left Direction = iota
	Right
	Up
	Down
)

var bodyIcon *ebiten.Image

type Snake struct {
	Body      []Point
	Direction Direction
	Speed     int
	Icon      *ebiten.Image
}

func init() {
	var err error
	bodyIcon, _, err = ebitenutil.NewImageFromFile("assets/body.png")
	if err != nil {
		log.Fatal(err)
	}
}

func NewSnake() *Snake {
	newBody := make([]Point, 0)
	newBody = append(newBody, Point{200, 360})

	return &Snake{
		Body:      newBody,
		Direction: Right,
		Speed:     0,
		Icon:      bodyIcon,
	}
}

func (s *Snake) Move() {
	xMove := 0
	yMove := 0
	if s.Direction == Up {
		xMove = 0
		yMove = -SQUARE_SIZE
	} else if s.Direction == Down {
		xMove = 0
		yMove = SQUARE_SIZE
	} else if s.Direction == Left {
		xMove = -SQUARE_SIZE
		yMove = 0
	} else if s.Direction == Right {
		xMove = SQUARE_SIZE
		yMove = 0
	}

	prevHead := s.Body[0]

	for i := len(s.Body) - 1; i > 0; i-- {
		s.Body[i].X, s.Body[i].Y = s.Body[i-1].X, s.Body[i-1].Y
	}

	s.Body[0].X, s.Body[0].Y = prevHead.X+xMove, prevHead.Y+yMove
}

func (s *Snake) ConsumeFood() {
	newHead := s.Body[0]

	if s.Direction == Up {
		newHead.X += 0
		newHead.Y += -SQUARE_SIZE
	} else if s.Direction == Down {
		newHead.X += 0
		newHead.Y += SQUARE_SIZE
	} else if s.Direction == Left {
		newHead.X += -SQUARE_SIZE
		newHead.Y += 0
	} else if s.Direction == Right {
		newHead.X += SQUARE_SIZE
		newHead.Y += 0
	}

	s.Body = append([]Point{newHead}, s.Body...)
}
