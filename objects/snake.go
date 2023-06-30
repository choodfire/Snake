package objects

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
	Speed     int
}

func NewSnake() *Snake {
	newBody := make([]Point, 0, 10)
	newBody = append(newBody, Point{LEFT_BORDER + ((RIGHT_BORDER - LEFT_BORDER) / 4),
		UPPER_BORDER + ((BOTTOM_BORDER - UPPER_BORDER) / 2)})

	return &Snake{
		Body:      newBody,
		Direction: Right,
		Speed:     0,
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
