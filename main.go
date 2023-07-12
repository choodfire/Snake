package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"log"
	"snake/game"
	"snake/objects"
)

func main() {
	ebiten.SetWindowSize(objects.SCREEN_WIDTH, objects.SCREEN_HEIGHT)
	ebiten.SetWindowTitle("Snake")
	if err := ebiten.RunGame(game.NewGame()); err != nil {
		log.Fatal(err)
	}
}
