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
	debug      bool
}

func NewGameState(len int) *GameState {
	return &GameState{
		gridOld:    *NewSeed(len),
		gridNew:    *NewGrid(len),
		length:     len,
		alives:     make([][]int, len, len*len),
		generation: 1,
		debug:      false,
	}
}

func (gs *GameState) Play() {
	for {
		// Initial
		gs.SetAlives(gs.GetOldGrid().GetAlivesPos())
		// Display
		gs.show()

		// Prepare next iteration
		gs.update()
		gs.transfertOldToNextGrid()
		gs.prepareNextIteration()
		time.Sleep(1 * time.Second)
	}
}

// --------------------------------------------

func (gs *GameState) update() {
	if gs.GetDebug() {
		fmt.Printf("DEBUG - Before UpdateCells:\n")
		fmt.Printf("  OldGrid population: %d\n", len(gs.GetOldGrid().GetAlivesPos()))
		fmt.Printf("  NewGrid population: %d\n", len(gs.GetNextGrid().GetAlivesPos()))
	}

	// Update new cells
	gs.SetAlives(gs.gridNew.UpdateCells(&gs.gridOld))

	if gs.GetDebug() {
		fmt.Printf("DEBUG - After UpdateCells:\n")
		fmt.Printf("  OldGrid population: %d\n", len(gs.GetOldGrid().GetAlivesPos()))
		fmt.Printf("  NewGrid population: %d\n", len(gs.GetNextGrid().GetAlivesPos()))
	}
}

func (gs *GameState) prepareNextIteration() {
	gs.updateGenerationNumber()
}

func (gs *GameState) transfertOldToNextGrid() {

	if gs.GetDebug() {
		fmt.Printf("DEBUG TRANSFERT - Before:\n")
		fmt.Printf("  oldGrid pop: %d\n", len(gs.gridOld.GetAlivesPos()))
		fmt.Printf("  newGrid pop: %d\n", len(gs.gridNew.GetAlivesPos()))
	}

	gs.gridOld = gs.gridNew
	gs.gridNew = *NewGrid(gs.GetLength())

	if gs.GetDebug() {
		fmt.Printf("DEBUG TRANSFERT - After:\n")
		fmt.Printf("  oldGrid pop: %d\n", len(gs.gridOld.GetAlivesPos()))
		fmt.Printf("  newGrid pop: %d\n", len(gs.gridNew.GetAlivesPos()))
	}
	fmt.Printf("---\n")
}

// --------------------------------------------

func (gs *GameState) show() {
	gs.showHeader()
	gs.GetOldGrid().Show()
}

func (gs *GameState) showHeader() {
	fmt.Println(color.Purple(), "------------------", color.Reset())
	fmt.Println(color.Cyan(), "  Generation:", gs.GetGeneration(), color.Reset())
	fmt.Println(color.Cyan(), "  Population:", len(*gs.GetAlives()), color.Reset())
	fmt.Println(color.Purple(), "------------------", color.Reset())
}

// --------------------------------------------

func (gs *GameState) GetOldGrid() *Grid {
	return &gs.gridOld
}

func (gs *GameState) GetNextGrid() *Grid {
	return &gs.gridNew
}

func (gs *GameState) GetLength() int {
	return gs.length
}

func (gs *GameState) GetAlives() *[][]int {
	return &gs.alives
}

func (gs *GameState) GetDebug() bool {
	return gs.debug
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
