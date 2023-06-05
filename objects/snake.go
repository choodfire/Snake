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
	Length    int
}

func NewSnake() *Snake {
	return &Snake{
		Body:      []Point{{100, 300}},
		Direction: Right,
		Length:    1,
	}
}
