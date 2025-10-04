package game

import (
	"fmt"
	"math/rand"
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

func NewSeed(n int) *Grid {
	m := *NewGrid(n)
	for range rand.Intn(n * n) {
		m[rand.Intn(n)][rand.Intn(n)].SetAlive(true)
	}
	return &m
}

func (gs *Game) Play() {
	// Game loop
	for {
		// Initial
		gs.intialize()
		// Display
		gs.display()
		// Prepare next iteration
		gs.prepareNextIteration()
	}
}

// --------------------------------------------

func (gs *Game) intialize() {
	gs.SetAlives(gs.GetOldGrid().GetAlives())
}

func (gs *Game) prepareNextIteration() {
	gs.update()
	gs.transfertOldToNextGrid()
	gs.updateGenerationNumber()
	// Game speed
	time.Sleep(time.Duration(gs.GetLag()) * time.Millisecond)
}

func (gs *Game) update() {
	// Update new cells
	gs.newGrid.UpdateCells(gs.GetOldGrid())
}

func (gs *Game) transfertOldToNextGrid() {
	gs.oldGrid = *gs.GetNextGrid()
	gs.newGrid = *NewGrid(gs.GetSize())
}

// --------------------------------------------

func (gs *Game) display() {
	clearDisplay()
	gs.displayHeader()
	gs.GetOldGrid().Display()
	fmt.Printf("Press [Ctrl + C] to stop.\n")
}

func clearDisplay() {
	fmt.Print("\033[H\033[2J")
}

func (gs *Game) displayHeader() {
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
