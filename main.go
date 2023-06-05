package main

import (
	"github.com/hajimehoshi/ebiten"
	"log"
	"snake/game"
)

func main() {
	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowTitle("Snake")
	if err := ebiten.RunGame(game.NewGame()); err != nil {
		log.Fatal(err)
	}
}
