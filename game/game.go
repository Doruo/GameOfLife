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

type Game struct {
	previousGrid Grid // Previous gen<eration
	nextGrid     Grid // New generation
	size         int
	alives       [][]int // Alives cells
	generation   int     // Generation number
	lag          int     // Lag frame/milliseconds
	debug        bool
}

func NewGame(size int) *Game {
	return &Game{
		previousGrid: *NewSeed(size),
		nextGrid:     *NewGrid(size),
		size:         size,
		alives:       make([][]int, size, size*size),
		generation:   1,
		lag:          300,
		debug:        false,
	}
}

func (gs *Game) Play() {
	// Game loop
	for {
		// Update
		gs.update()
		// Prepare next iteration
		gs.prepareNextIteration()
	}
}

// --------------------------------------------

func (gs *Game) update() {
	clearDisplay()
	gs.displayHeader()
	// Updates and display next grid state with all logical process
	gs.SetAlives(gs.nextGrid.UpdateCells(gs.GetPreviousGrid()))
	fmt.Printf("Press [Ctrl + C] to stop.\n")
}

func (gs *Game) prepareNextIteration() {
	gs.transfertPreviousToNextGrid()
	gs.updateGeneration()
	// Game speed
	time.Sleep(time.Duration(gs.GetLag()) * time.Millisecond)
}

func (gs *Game) transfertPreviousToNextGrid() {
	gs.previousGrid = *gs.GetNextGrid()
	gs.nextGrid = *NewGrid(gs.GetSize())
}

// --------------------------------------------

func clearDisplay() {
	fmt.Print("\033[H\033[2J")
}

func (gs *Game) displayHeader() {
	fmt.Println(Purple(), "------------------", ColorReset())
	fmt.Println(Cyan(), "  Generation:", gs.GetGeneration(), ColorReset())
	fmt.Println(Cyan(), "  Population:", len(*gs.GetAlives()), ColorReset())
	fmt.Println(Purple(), "------------------", ColorReset())
}

// --------------------------------------------

func (gs *Game) GetPreviousGrid() *Grid {
	return &gs.previousGrid
}

func (gs *Game) GetNextGrid() *Grid {
	return &gs.nextGrid
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

func (gs *Game) updateGeneration() {
	gs.generation++
}
