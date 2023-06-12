package game

import (
	"errors"
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"
	"github.com/hajimehoshi/ebiten/text"
	"golang.org/x/image/font/basicfont"
	"image/color"
	_ "image/png"
	"snake/objects"
)

type Game struct {
	food      *objects.Food
	snake     *objects.Snake
	isRunning bool
	isPaused  bool
	speed     int
	maxSpeed  int
}

func NewGame() *Game {
	return &Game{
		food:      objects.NewFood(),
		snake:     objects.NewSnake(),
		isRunning: true,
		isPaused:  false,
		speed:     0,
		maxSpeed:  10,
	}
}

func (g *Game) Update(screen *ebiten.Image) error {
	if g.isRunning == false {
		if inpututil.IsKeyJustPressed(ebiten.KeyR) {
			g.Restart()
		}
		return nil
	}

	if g.isPaused == true {
		if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
			g.isPaused = false
		}
		return nil
	}

	if g.speed < g.maxSpeed {
		g.speed += 1
		return nil
	}
	g.speed = 0

	if (ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeyUp)) && g.snake.Direction != objects.Down { // maybe method
		g.snake.Direction = objects.Up
	} else if (ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyLeft)) && g.snake.Direction != objects.Right {
		g.snake.Direction = objects.Left
	} else if (ebiten.IsKeyPressed(ebiten.KeyS) || ebiten.IsKeyPressed(ebiten.KeyDown)) && g.snake.Direction != objects.Up {
		g.snake.Direction = objects.Down
	} else if (ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyRight)) && g.snake.Direction != objects.Left {
		g.snake.Direction = objects.Right
	} else if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		g.isPaused = true
	}

	if g.CheckGameOver() != nil {
		g.isRunning = false
		return nil
	}

	g.snake.Move()

	if head := g.snake.Body[0]; head == g.food.Point {
		g.snake.ConsumeFood()
		g.SpawnNewFood()
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{R: 0, G: 0, B: 0, A: 255})

	if g.isRunning == false {
		g.GameOverScreen(screen)
	}

	if g.isPaused == true {
		g.GamePausedScreen(screen)
	}

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

	if head.X > objects.SCREEN_WIDTH-objects.SQUARE_SIZE-1 && g.snake.Direction == objects.Right {
		return errors.New("snake hit right border")
	}

	if head.X < objects.SQUARE_SIZE-1 && g.snake.Direction == objects.Left {
		return errors.New("snake hit left border")
	}

	if head.Y > objects.SCREEN_HEIGHT-objects.SQUARE_SIZE-1 && g.snake.Direction == objects.Down {
		return errors.New("snake hit bottom border")
	}

	if head.Y < objects.SQUARE_SIZE-1 && g.snake.Direction == objects.Up {
		return errors.New("snake hit upper border")
	}

	for _, bodyPoint := range g.snake.Body[1:] {
		if head == bodyPoint {
			return errors.New("snake hit itself")
		}
	}

	return nil
}

func (g *Game) Restart() {
	g.food = objects.NewFood()
	g.snake = objects.NewSnake()
	g.isRunning = true
	g.isPaused = false
	g.speed = 0
	g.maxSpeed = 10
}

func (g *Game) GameOverScreen(screen *ebiten.Image) {
	cx := objects.SCREEN_WIDTH / 2
	cy := objects.SCREEN_HEIGHT / 2
	face := basicfont.Face7x13

	scoreText := fmt.Sprintf("Your score: %d.", len(g.snake.Body))
	scoreBounds := text.BoundString(face, scoreText)
	scoreTextX, scoreTextY := cx-scoreBounds.Min.X-scoreBounds.Dx()/2, cy-scoreBounds.Min.Y-scoreBounds.Dy()/2
	text.Draw(screen, scoreText, face, scoreTextX, scoreTextY-10, objects.WHITE)

	restartText := "Game over. Press \"R\" to restart."
	restartBounds := text.BoundString(face, restartText)
	restartTextX, restartTextY := cx-restartBounds.Min.X-restartBounds.Dx()/2, cy-restartBounds.Min.Y-restartBounds.Dy()/2
	text.Draw(screen, restartText, face, restartTextX, restartTextY+10, objects.WHITE)
}

func (g *Game) GamePausedScreen(screen *ebiten.Image) {
	face := basicfont.Face7x13
	restartText := "Game paused. Press \"Esc\" to continue."
	bounds := text.BoundString(face, restartText)
	cx := objects.SCREEN_WIDTH / 2
	cy := objects.SCREEN_HEIGHT / 2
	x, y := cx-bounds.Min.X-bounds.Dx()/2, cy-bounds.Min.Y-bounds.Dy()/2

	text.Draw(screen, restartText, face, x, y, objects.WHITE)
}

func (g *Game) SpawnNewFood() {
	g.food = objects.NewFood()

	for _, bodyPoint := range g.snake.Body {
		if g.food.Point == bodyPoint {
			g.SpawnNewFood()
			break
		}
	}
}
