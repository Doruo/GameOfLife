package game

import (
	"fmt"
	"time"

	"github.com/doruo/gameoflife/game/color"
)

type GameState struct {
	grid   Grid
	length int
	alives [length][length]int
}

func NewGameState(len int) *GameState {
	return &GameState{
		grid:   *NewSeed(len),
		length: len,
	}
}

func (gs *GameState) Play() {
	i := 0
	for {
		fmt.Println(color.Purple, "------------------", color.Reset)
		fmt.Println(color.Purple, "  Generation:", i, color.Reset)
		fmt.Println(color.Purple, "------------------", color.Reset)
		gs.grid.Show()
		time.Sleep(1 * time.Second)
		gs.grid.UpdateCells()
		i++
	}
}

func (gs *GameState) GetGrid() Grid {
	return gs.grid
}

func (gs *GameState) GetLength() int {
	return gs.length
}
