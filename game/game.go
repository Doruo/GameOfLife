package game

import (
	"fmt"
	"time"

	"github.com/doruo/gameoflife/game/color"
)

type GameState struct {
	gridOld    Grid // Previous generation
	gridNew    Grid // New generation
	length     int
	alives     [][]int
	generation int
}

func NewGameState(len int) *GameState {
	return &GameState{
		gridOld:    *NewSeed(len),
		gridNew:    *NewGrid(len),
		length:     len,
		alives:     make([][]int, len, len*len),
		generation: 1,
	}
}

func (gs *GameState) Play() {
	for {
		// Process
		gs.update()
		// Display
		gs.show()
		// Prepare next iteration
		gs.prepareNextIteration()
		time.Sleep(1 * time.Second)
	}
}

func (gs *GameState) update() {
	gs.GetNextGrid().UpdateCells(gs.GetOldGrid())
	gs.SetAlives(gs.GetNextGrid().GetAlivesPos())
}

func (gs *GameState) prepareNextIteration() {
	gs.updateGenerationNumber()
	gs.transfertOldToNextGrid()
}

func (gs *GameState) show() {
	gs.showHeader()
	gs.GetNextGrid().Show()
}

func (gs *GameState) showHeader() {
	fmt.Println(color.Purple(), "------------------", color.Reset())
	fmt.Println(color.Cyan(), "  Generation:", gs.GetGeneration(), color.Reset())
	fmt.Println(color.Cyan(), "  Population:", len(gs.GetAlives()), color.Reset())
	fmt.Println(color.Purple(), "------------------", color.Reset())
}

func (gs *GameState) transfertOldToNextGrid() {
	gs.gridOld = gs.gridNew
	gs.gridNew = *NewGrid(len(gs.gridOld))
}

func (gs *GameState) GetOldGrid() Grid {
	return gs.gridOld
}

func (gs *GameState) GetNextGrid() Grid {
	return gs.gridNew
}

func (gs *GameState) GetLength() int {
	return gs.length
}

func (gs *GameState) GetAlives() [][]int {
	return gs.alives
}

func (gs *GameState) SetAlives(alives [][]int) {
	gs.alives = alives
}

func (gs *GameState) GetGeneration() int {
	return gs.generation
}

func (gs *GameState) updateGenerationNumber() {
	gs.generation++
}
