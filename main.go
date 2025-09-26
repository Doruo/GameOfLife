package main

import "github.com/doruo/gameoflife/game"

func main() {
	game.Show(*game.NewSeed(3))
}
