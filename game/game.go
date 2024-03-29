package game

import (
	"errors"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font/basicfont"
	_ "image/png"
	"log"
	"snake/objects"
)

type Game struct {
	food              *objects.Food
	snake             *objects.Snake
	isRunning         bool
	isPaused          bool
	isEducationScreen bool
	maxSnakeSpeed     int
	currentSpeed      int
	currentScore      int
	maxScore          int
	background        *ebiten.Image
	educationScreen   *ebiten.Image
}

func NewGame() *Game {
	img, _, err := ebitenutil.NewImageFromFile("assets/background.png")
	if err != nil {
		log.Fatal(err)
	}

	educationScreen, _, err := ebitenutil.NewImageFromFile("assets/education.png")
	if err != nil {
		log.Fatal(err)
	}

	return &Game{
		food:              objects.NewFood(),
		snake:             objects.NewSnake(),
		isRunning:         true,
		isPaused:          false,
		isEducationScreen: true,
		maxSnakeSpeed:     10,
		currentSpeed:      0,
		currentScore:      0,
		maxScore:          0,
		background:        img,
		educationScreen:   educationScreen,
	}
}

func (g *Game) Update() error {
	if g.isEducationScreen == true {
		if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
			g.isEducationScreen = false
		}
		return nil
	}

	if g.isRunning == false {
		if inpututil.IsKeyJustPressed(ebiten.KeyR) {
			g.Restart()
		}
		return nil
	}

	if g.isPaused == true {
		if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
			g.isPaused = false
		}
		return nil
	}

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

	if g.currentSpeed < g.maxSnakeSpeed {
		g.currentSpeed += 1
		return nil
	}
	g.currentSpeed = g.snake.Speed

	if g.CheckGameOver() != nil {
		g.isRunning = false
		return nil
	}

	g.snake.Move()

	if head := g.snake.Body[0]; head == g.food.Point {
		g.snake.ConsumeFood()
		g.SpawnNewFood()
		g.currentScore += g.food.Score
		if g.currentScore > g.maxScore {
			g.maxScore += g.food.Score
		}

		g.SpeedUpSnake()
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(g.background, nil)

	if g.isEducationScreen == true {
		g.FirstScreen(screen)
		return
	}

	if g.isRunning == false {
		g.GameOverScreen(screen)
	}

	if g.isPaused == true {
		g.GamePausedScreen(screen)
	}

	g.DrawScoreText(screen)

	options := &ebiten.DrawImageOptions{}
	options.GeoM.Translate(float64(g.food.Point.X), float64(g.food.Point.Y))
	screen.DrawImage(g.food.Icon, options)

	for _, point := range g.snake.Body {
		options := &ebiten.DrawImageOptions{}
		options.GeoM.Translate(float64(point.X), float64(point.Y))
		screen.DrawImage(g.snake.Icon, options)
	}
}

func (g *Game) Layout(outsideWidth int, outsideHeight int) (screenWidth int, screenHeight int) {
	return objects.SCREEN_WIDTH, objects.SCREEN_HEIGHT
}

func (g *Game) CheckGameOver() error {
	head := g.snake.Body[0]

	if (head.X > objects.RIGHT_BORDER-objects.SQUARE_SIZE-1 && g.snake.Direction == objects.Right) ||
		head.X > objects.RIGHT_BORDER-objects.SQUARE_SIZE {
		return errors.New("snake hit right border")
	}

	if (head.X < objects.LEFT_BORDER+objects.SQUARE_SIZE+1 && g.snake.Direction == objects.Left) ||
		head.X < objects.LEFT_BORDER+objects.SQUARE_SIZE {
		return errors.New("snake hit left border")
	}

	if (head.Y > objects.BOTTOM_BORDER-objects.SQUARE_SIZE-1 && g.snake.Direction == objects.Down) ||
		head.X > objects.BOTTOM_BORDER-objects.SQUARE_SIZE {
		return errors.New("snake hit bottom border")
	}

	if (head.Y < objects.UPPER_BORDER+objects.SQUARE_SIZE+1 && g.snake.Direction == objects.Up) ||
		head.X < objects.UPPER_BORDER+objects.SQUARE_SIZE {
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
	g.maxSnakeSpeed = 10
	g.currentSpeed = 0
	g.currentScore = 0
}

func (g *Game) GameOverScreen(screen *ebiten.Image) {
	cx := objects.SCREEN_WIDTH / 2
	cy := objects.SCREEN_HEIGHT / 2
	face := basicfont.Face7x13

	scoreText := fmt.Sprintf("Your score: %d.", g.currentScore)
	scoreBounds := text.BoundString(face, scoreText)
	scoreTextX, scoreTextY := cx-scoreBounds.Min.X-scoreBounds.Dx()/2, cy-scoreBounds.Min.Y-scoreBounds.Dy()/2
	text.Draw(screen, scoreText, face, scoreTextX, scoreTextY-10, colornames.Black)

	restartText := "Game over. Press \"R\" to restart."
	restartBounds := text.BoundString(face, restartText)
	restartTextX, restartTextY := cx-restartBounds.Min.X-restartBounds.Dx()/2, cy-restartBounds.Min.Y-restartBounds.Dy()/2
	text.Draw(screen, restartText, face, restartTextX, restartTextY+10, colornames.Black)
}

func (g *Game) GamePausedScreen(screen *ebiten.Image) {
	face := basicfont.Face7x13
	restartText := "Game paused. Press \"Enter\" to continue."
	bounds := text.BoundString(face, restartText)
	cx := objects.SCREEN_WIDTH / 2
	cy := objects.SCREEN_HEIGHT / 2
	x, y := cx-bounds.Min.X-bounds.Dx()/2, cy-bounds.Min.Y-bounds.Dy()/2

	text.Draw(screen, restartText, face, x, y, colornames.Black)
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

func (g *Game) SpeedUpSnake() {
	if g.currentScore%4 == 0 && g.snake.Speed < g.maxSnakeSpeed {
		g.snake.Speed += 1
		g.currentSpeed += 1
	}
}

func (g *Game) DrawScoreText(screen *ebiten.Image) {
	face := basicfont.Face7x13
	currentScoreText := fmt.Sprintf("Current score: %d", g.currentScore)
	bounds := text.BoundString(face, currentScoreText)
	cx := objects.SCREEN_WIDTH / 4
	cy := objects.UPPER_BORDER / 2
	x, y := cx-bounds.Min.X-bounds.Dx()/2, cy-bounds.Min.Y-bounds.Dy()/2

	text.Draw(screen, currentScoreText, face, x, y, colornames.Black)

	bestScoreText := fmt.Sprintf("Best score: %d", g.maxScore)
	bounds = text.BoundString(face, currentScoreText)
	cx = objects.SCREEN_WIDTH / 4 * 3
	cy = objects.UPPER_BORDER / 2
	x, y = cx-bounds.Min.X-bounds.Dx()/2, cy-bounds.Min.Y-bounds.Dy()/2

	text.Draw(screen, bestScoreText, face, x, y, colornames.Black)
}

func (g *Game) StartGame() {
	ebiten.SetWindowSize(objects.SCREEN_WIDTH, objects.SCREEN_HEIGHT)
	ebiten.SetWindowTitle("Snake")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}

func (g *Game) FirstScreen(screen *ebiten.Image) {
	screen.DrawImage(g.educationScreen, nil)
}
