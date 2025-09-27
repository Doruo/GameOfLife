package game

import (
	"fmt"
	"time"

	"github.com/doruo/gameoflife/game/color"
)

type Game struct {
	oldGrid    Grid // Previous generation
	newGrid    Grid // New generation
	size       int
	alives     [][]int
	generation int
	lag        int // Lag frame/milliseconds
	debug      bool
}

func NewGame(size int) *Game {
	return &Game{
		oldGrid:    *NewSeed(size),
		newGrid:    *NewGrid(size),
		size:       size,
		alives:     make([][]int, size, size*size),
		generation: 1,
		lag:        500,
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
		time.Sleep(time.Duration(gs.GetLag()) * time.Millisecond)
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
	gs.SetAlives(gs.newGrid.UpdateCells(&gs.oldGrid))

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
		fmt.Printf("  oldGrid pop: %d\n", len(gs.oldGrid.GetAlivesPos()))
		fmt.Printf("  newGrid pop: %d\n", len(gs.newGrid.GetAlivesPos()))
	}

	gs.oldGrid = gs.newGrid
	gs.newGrid = *NewGrid(gs.GetSize())

	if gs.GetDebug() {
		fmt.Printf("DEBUG TRANSFERT - After:\n")
		fmt.Printf("  oldGrid pop: %d\n", len(gs.oldGrid.GetAlivesPos()))
		fmt.Printf("  newGrid pop: %d\n", len(gs.newGrid.GetAlivesPos()))
	}
	fmt.Printf("Press [Ctrl + C] to stop.\n")
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
	return &gs.oldGrid
}

func (gs *Game) GetNextGrid() *Grid {
	return &gs.newGrid
}

func (gs *Game) GetSize() int {
	return gs.size
}

func (gs *Game) GetAlives() *[][]int {
	return &gs.alives
}

func (gs *Game) GetLag() int {
	return gs.lag
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
