package game

import "github.com/hajimehoshi/ebiten"

type Game struct {
}

func NewGame() *Game {
	return &Game{}
}

func (g *Game) Update(screen *ebiten.Image) error {
	return nil
}

func (g *Game) Draw() {

}

func (g *Game) Layout(outsideWidth int, outsideHeight int) (screenWidth int, screenHeight int) {
	return 800, 600
}
