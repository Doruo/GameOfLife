package game

import (
	"fmt"
	"time"

	"github.com/doruo/gameoflife/game/color"
)

type Game struct {
	gridOld    Grid // Previous generation
	gridNew    Grid // New generation
	size       int
	alives     [][]int
	generation int
	debug      bool
}

func NewGame(size int) *Game {
	return &Game{
		gridOld:    *NewSeed(size),
		gridNew:    *NewGrid(size),
		size:       size,
		alives:     make([][]int, size, size*size),
		generation: 1,
		debug:      false,
	}
}

func (gs *Game) Play() {
	// Game loop
	for {
		// Initial
		gs.intialize()
		// Display
		gs.show()

		// Prepare next iteration
		gs.update()
		gs.prepareNextIteration()
		// Game speed
		time.Sleep(1 * time.Second)
	}
}

func (gs *Game) intialize() {
	gs.SetAlives(gs.GetOldGrid().GetAlivesPos())
}

// --------------------------------------------

func (gs *Game) update() {
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

func (gs *Game) prepareNextIteration() {
	gs.transfertOldToNextGrid()
	gs.updateGenerationNumber()
}

func (gs *Game) transfertOldToNextGrid() {

	if gs.GetDebug() {
		fmt.Printf("DEBUG TRANSFERT - Before:\n")
		fmt.Printf("  oldGrid pop: %d\n", len(gs.gridOld.GetAlivesPos()))
		fmt.Printf("  newGrid pop: %d\n", len(gs.gridNew.GetAlivesPos()))
	}

	gs.gridOld = gs.gridNew
	gs.gridNew = *NewGrid(gs.GetSize())

	if gs.GetDebug() {
		fmt.Printf("DEBUG TRANSFERT - After:\n")
		fmt.Printf("  oldGrid pop: %d\n", len(gs.gridOld.GetAlivesPos()))
		fmt.Printf("  newGrid pop: %d\n", len(gs.gridNew.GetAlivesPos()))
	}
	fmt.Printf("---\n")
}

// --------------------------------------------

func clearDisplay() {
	fmt.Print("\033[H\033[2J")
}

func (gs *Game) show() {
	clearDisplay()
	gs.showHeader()
	gs.GetOldGrid().Show()
}

func (gs *Game) showHeader() {
	fmt.Println(color.Purple(), "------------------", color.Reset())
	fmt.Println(color.Cyan(), "  Generation:", gs.GetGeneration(), color.Reset())
	fmt.Println(color.Cyan(), "  Population:", len(*gs.GetAlives()), color.Reset())
	fmt.Println(color.Purple(), "------------------", color.Reset())
}

// --------------------------------------------

func (gs *Game) GetOldGrid() *Grid {
	return &gs.gridOld
}

func (gs *Game) GetNextGrid() *Grid {
	return &gs.gridNew
}

func (gs *Game) GetSize() int {
	return gs.size
}

func (gs *Game) GetAlives() *[][]int {
	return &gs.alives
}

func (gs *Game) GetDebug() bool {
	return gs.debug
}

func (gs *Game) SetAlives(alives [][]int) {
	gs.alives = alives
}

func (gs *Game) GetGeneration() int {
	return gs.generation
}

func (gs *Game) updateGenerationNumber() {
	gs.generation++
}
