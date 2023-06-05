package game

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"image/color"
	_ "image/png"
	"snake/objects"
)

type Game struct {
	food  *objects.Food
	snake *objects.Snake

	running bool
}

func NewGame() *Game {
	return &Game{
		food:    objects.NewFood(),
		snake:   objects.NewSnake(),
		running: true,
	}
}

func (g *Game) Update(screen *ebiten.Image) error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0, 0, 0, 255})

	for _, point := range g.snake.Body {
		ebitenutil.DrawRect(screen, float64(point.X), float64(point.Y), 10, 10, color.RGBA{0, 255, 0, 255})
	}
}

func (g *Game) Layout(outsideWidth int, outsideHeight int) (screenWidth int, screenHeight int) {
	return 800, 800
}
