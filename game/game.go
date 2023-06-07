package game

import (
	"fmt"
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

const (
	SCREEN_WIDTH  = 800
	SCREEN_HEIGHT = 600
	SQUARE_SIZE   = 10
)

type Game struct {
	food          *objects.Food
	snake         *objects.Snake
	running       bool
	speed         int
	updateCounter int
}

func NewGame() *Game {
	return &Game{
		food:          objects.NewFood(),
		snake:         objects.NewSnake(),
		running:       true,
		speed:         10,
		updateCounter: 0,
	}
}

func (g *Game) Update(screen *ebiten.Image) error {
	if g.updateCounter < g.speed { // rethink
		g.updateCounter += 1
		return nil
	}
	g.updateCounter = 0

	if g.running == false {
		fmt.Println("Game over!")
		// todo: stop the game
	}

	g.snake.Move()

	if (ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeyUp)) && g.snake.Direction != objects.Down { // maybe method
		g.snake.Direction = objects.Up
	} else if (ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyLeft)) && g.snake.Direction != objects.Right {
		g.snake.Direction = objects.Left
	} else if (ebiten.IsKeyPressed(ebiten.KeyS) || ebiten.IsKeyPressed(ebiten.KeyDown)) && g.snake.Direction != objects.Up {
		g.snake.Direction = objects.Down
	} else if (ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyRight)) && g.snake.Direction != objects.Left {
		g.snake.Direction = objects.Right
	}

	head := g.snake.Body[0]

	if head == g.food.Point {
		g.snake.ConsumeFood()
		g.food = objects.NewFood()
	}

	if g.snake.CheckBorders() != nil { // questionable
		g.running = false
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{R: 0, G: 0, B: 0, A: 255})

	ebitenutil.DrawRect(screen, float64(g.food.Point.X), float64(g.food.Point.Y), SQUARE_SIZE, SQUARE_SIZE, RED)

	for _, point := range g.snake.Body {
		ebitenutil.DrawRect(screen, float64(point.X), float64(point.Y), SQUARE_SIZE, SQUARE_SIZE, GREEN)
	}
}

func (g *Game) Layout(outsideWidth int, outsideHeight int) (screenWidth int, screenHeight int) {
	return SCREEN_WIDTH, SCREEN_HEIGHT
}
