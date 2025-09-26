package main

import (
	"fmt"
	"time"

	"github.com/doruo/gameoflife/game"
)

func main() {
	g := game.NewGameState(3).Grid

	for {
		g.UpdateCells()
		g.Show()
		time.Sleep(1 * time.Second)
		fmt.Println(" ")
	}
}
