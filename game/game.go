package game

import (
	"fmt"
	"time"
)

func ColorReset() string { return "\033[0m" }
func Red() string        { return "\033[31m" }
func Green() string      { return "\033[32m" }
func Purple() string     { return "\033[35m" }
func Cyan() string       { return "\033[36m" }

type GameOfLife struct {
	previousGrid Grid // Previous generation
	nextGrid     Grid // New generation
	size         int
	alives       [][]int // Alives cells
	generation   int     // Generation number
	lag          int     // Lag frame/milliseconds
	debug        bool
}

func NewGameOfLife(size int) *GameOfLife {
	return &GameOfLife{
		previousGrid: *NewSeed(size),
		nextGrid:     *NewGrid(size),
		size:         size,
		alives:       make([][]int, size, size*size),
		generation:   1,
		lag:          300,
		debug:        false,
	}
}

func (gs *GameOfLife) Play() {
	// Game loop
	for {
		// Update
		gs.update()
		// Prepare next iteration
		gs.prepareNextIteration()
	}
}

// --------------------------------------------

func (gs *GameOfLife) update() {
	clearDisplay()
	gs.displayHeader()
	// Updates and display next grid state with all logical process
	gs.SetAlives(gs.nextGrid.UpdateCells(gs.GetPreviousGrid()))
	fmt.Printf("Press [Ctrl + C] to stop.\n")
}

func (gs *GameOfLife) prepareNextIteration() {
	gs.transfertPreviousToNextGrid()
	gs.updateGeneration()
	// Game speed
	time.Sleep(time.Duration(gs.GetLag()) * time.Millisecond)
}

func (gs *GameOfLife) transfertPreviousToNextGrid() {
	gs.previousGrid = *gs.GetNextGrid()
	gs.nextGrid = *NewGrid(gs.GetSize())
}

// --------------------------------------------

func clearDisplay() {
	fmt.Print("\033[H\033[2J")
}

func (gs *GameOfLife) displayHeader() {
	fmt.Println(Purple(), "------------------", ColorReset())
	fmt.Println(Cyan(), "  Generation:", gs.GetGeneration(), ColorReset())
	fmt.Println(Cyan(), "  Population:", len(*gs.GetAlives()), ColorReset())
	fmt.Println(Purple(), "------------------", ColorReset())
}

// --------------------------------------------

func (gs *GameOfLife) GetPreviousGrid() *Grid {
	return &gs.previousGrid
}

func (gs *GameOfLife) GetNextGrid() *Grid {
	return &gs.nextGrid
}

func (gs *GameOfLife) GetSize() int {
	return gs.size
}

func (gs *GameOfLife) GetAlives() *[][]int {
	return &gs.alives
}

func (gs *GameOfLife) GetLag() int {
	return gs.lag
}

func (gs *GameOfLife) GetDebug() bool {
	return gs.debug
}

func (gs *GameOfLife) SetAlives(alives [][]int) {
	gs.alives = alives
}

func (gs *GameOfLife) GetGeneration() int {
	return gs.generation
}

func (gs *GameOfLife) updateGeneration() {
	gs.generation++
}
