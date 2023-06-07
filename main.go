package main

import (
	"github.com/hajimehoshi/ebiten"
	"log"
	"snake/game"
)

func main() {
	ebiten.SetWindowSize(game.SCREEN_WIDTH, game.SCREEN_HEIGHT)
	ebiten.SetWindowTitle("Snake")
	if err := ebiten.RunGame(game.NewGame()); err != nil {
		log.Fatal(err)
	}
}
