package objects

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
	"math/rand"
)

type Food struct {
	Icon  *ebiten.Image
	Score int
	Point Point
}

var food1Icon *ebiten.Image
var food2Icon *ebiten.Image
var food3Icon *ebiten.Image
var food4Icon *ebiten.Image

func init() {
	var err error
	food1Icon, _, err = ebitenutil.NewImageFromFile("assets/images/food1.png")
	if err != nil {
		log.Fatal(err)
	}
	food2Icon, _, err = ebitenutil.NewImageFromFile("assets/images/food2.png")
	if err != nil {
		log.Fatal(err)
	}
	food3Icon, _, err = ebitenutil.NewImageFromFile("assets/images/food3.png")
	if err != nil {
		log.Fatal(err)
	}
	food4Icon, _, err = ebitenutil.NewImageFromFile("assets/images/food4.png")
	if err != nil {
		log.Fatal(err)
	}
}

func NewFood() *Food {
	min := (LEFT_BORDER + SQUARE_SIZE) / SQUARE_SIZE
	max := (UPPER_BORDER + SQUARE_SIZE) / SQUARE_SIZE

	score := 1
	var icon *ebiten.Image

	randNum := rand.Float32()
	if randNum < 0.8 {
		icon = food1Icon
	} else if randNum < 0.9 {
		score = 2
		icon = food2Icon
	} else if randNum < 0.99 {
		score = 5
		icon = food3Icon
	} else {
		score = 10
		icon = food4Icon
	}

	return &Food{
		Point: Point{
			X: (rand.Intn((RIGHT_BORDER/SQUARE_SIZE)-min) + min) * SQUARE_SIZE,
			Y: (rand.Intn((BOTTOM_BORDER/SQUARE_SIZE)-max) + max) * SQUARE_SIZE,
		},
		Score: score,
		Icon:  icon,
	}
}
