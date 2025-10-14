package main

import (
	"github.com/doruo/gameoflife/gol"
)

func main() {
	width, heigth := 1920, 1080
	game := gol.NewGame(width, heigth)
	game.Run()
}
