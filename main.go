package main

import (
	"github.com/doruo/gameoflife/game"
)

func main() {
	game.NewGameState(15).Play()
}
