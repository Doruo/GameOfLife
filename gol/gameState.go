package gol

import (
	"fmt"
	"time"
)

func ColorReset() string { return "\033[0m" }
func Red() string        { return "\033[31m" }
func Green() string      { return "\033[32m" }
func Purple() string     { return "\033[35m" }
func Cyan() string       { return "\033[36m" }

type GameState struct {
	previousGrid          Grid    // Previous generation
	nextGrid              Grid    // New generation
	alives                [][]int // Alives cells
	size, generation, lag int     // Generation number, Lag frame/milliseconds
	debug                 bool
}

func NewGameState(size int) *GameState {
	return &GameState{
		previousGrid: *NewSeed(size),
		nextGrid:     *NewGrid(size),
		size:         size,
		alives:       make([][]int, size, size*size),
		generation:   1,
		lag:          300,
		debug:        false,
	}
}

func (gs *GameState) Play() {
	// Game stateloop
	for {
		// Update
		gs.update()
		// Prepare next iteration
		gs.prepareNextIteration()
	}
}

// --------------------------------------------

func (gs *GameState) update() {
	clearDisplay()
	gs.displayHeader()
	// Updates and display next grid state with all logical process
	gs.SetAlives(gs.nextGrid.UpdateCells(gs.GetPreviousGrid()))
	fmt.Printf("Press [Ctrl + C] to stop.\n")
}

func (gs *GameState) prepareNextIteration() {
	gs.transfertPreviousToNextGrid()
	gs.updateGeneration()
	// Game speed
	time.Sleep(time.Duration(gs.GetLag()) * time.Millisecond)
}

func (gs *GameState) transfertPreviousToNextGrid() {
	gs.previousGrid = *gs.GetNextGrid()
	gs.nextGrid = *NewGrid(gs.GetSize())
}

// --------------------------------------------

func clearDisplay() {
	fmt.Print("\033[H\033[2J")
}

func (gs *GameState) displayHeader() {
	fmt.Println(Purple(), "------------------", ColorReset())
	fmt.Println(Cyan(), "  Generation:", gs.GetGeneration(), ColorReset())
	fmt.Println(Cyan(), "  Population:", len(*gs.GetAlives()), ColorReset())
	fmt.Println(Purple(), "------------------", ColorReset())
}

// --------------------------------------------

func (gs *GameState) GetPreviousGrid() *Grid {
	return &gs.previousGrid
}

func (gs *GameState) GetNextGrid() *Grid {
	return &gs.nextGrid
}

func (gs *GameState) GetSize() int {
	return gs.size
}

func (gs *GameState) GetAlives() *[][]int {
	return &gs.alives
}

func (gs *GameState) GetLag() int {
	return gs.lag
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

func (gs *GameState) updateGeneration() {
	gs.generation++
}
