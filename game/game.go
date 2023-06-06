package game

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"image/color"
	_ "image/png"
	"snake/objects"
)

var (
	RED   = color.RGBA{R: 255, G: 0, B: 0, A: 255}
	GREEN = color.RGBA{R: 0, G: 255, B: 0, A: 255}
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
	g.snake.Move()

	if ebiten.IsKeyPressed(ebiten.KeyW) {
		g.snake.Direction = objects.Up
	} else if ebiten.IsKeyPressed(ebiten.KeyA) {
		g.snake.Direction = objects.Left
	} else if ebiten.IsKeyPressed(ebiten.KeyS) {
		g.snake.Direction = objects.Down
	} else if ebiten.IsKeyPressed(ebiten.KeyD) {
		g.snake.Direction = objects.Right
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{R: 0, G: 0, B: 0, A: 255})

	ebitenutil.DrawRect(screen, float64(g.food.Point.X), float64(g.food.Point.Y), 10, 10, RED)

	for _, point := range g.snake.Body {
		ebitenutil.DrawRect(screen, float64(point.X), float64(point.Y), 10, 10, GREEN)
	}
}

func (g *Game) Layout(outsideWidth int, outsideHeight int) (screenWidth int, screenHeight int) {
	return 800, 800
}
