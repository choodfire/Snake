package game

import (
	"errors"
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"image/color"
	_ "image/png"
	"snake/objects"
)

type Game struct {
	food     *objects.Food
	snake    *objects.Snake
	running  bool
	speed    int
	maxSpeed int
}

func NewGame() *Game {
	return &Game{
		food:     objects.NewFood(),
		snake:    objects.NewSnake(),
		running:  true,
		speed:    0,
		maxSpeed: 10,
	}
}

func (g *Game) Update(screen *ebiten.Image) error {
	if g.speed < g.maxSpeed {
		g.speed += 1
		return nil
	}
	g.speed = 0

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

	if g.CheckGameOver() != nil {
		g.running = false
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{R: 0, G: 0, B: 0, A: 255})

	ebitenutil.DrawRect(screen, float64(g.food.Point.X), float64(g.food.Point.Y), objects.SQUARE_SIZE, objects.SQUARE_SIZE, objects.RED)

	for _, point := range g.snake.Body {
		ebitenutil.DrawRect(screen, float64(point.X), float64(point.Y), objects.SQUARE_SIZE, objects.SQUARE_SIZE, objects.GREEN)
	}
}

func (g *Game) Layout(outsideWidth int, outsideHeight int) (screenWidth int, screenHeight int) {
	return objects.SCREEN_WIDTH, objects.SCREEN_HEIGHT
}

func (g *Game) CheckGameOver() error {
	head := g.snake.Body[0]

	if head.X > objects.SCREEN_WIDTH || head.X < 0 || head.Y > objects.SCREEN_HEIGHT || head.Y < 0 {
		return errors.New("snake hit border")
	}

	for _, bodyPoint := range g.snake.Body[1:] {
		if head == bodyPoint {
			return errors.New("snake itself")
		}
	}

	return nil
}
