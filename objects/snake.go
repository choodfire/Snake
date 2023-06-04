package objects

type Direction int

const (
	Left Direction = iota
	Right
	Up
	Down
)

type Snake struct {
	body      []Point
	direction Direction
	length    int
}
