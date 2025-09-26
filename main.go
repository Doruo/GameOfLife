package main

import (
	"fmt"
	"math/rand"
)

type Grid [][]bool

func newGrid() *Grid {
	return &Grid{}
}

func randomInt(limit int) (int, int) {
	return rand.Intn(limit), rand.Intn(limit)
}

func newSeed() *Grid {
	// limit := 10
	g := newGrid()
	for i := 0; i < rand.Int(); i++ {
		// g[randomInt(limit)][randomInt(limit)] == true
	}
	return g
}

func main() {
	fmt.Print("Game of Life")
	grid := newSeed()
	fmt.Print(grid)
}
