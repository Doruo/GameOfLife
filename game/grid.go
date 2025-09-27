package game

import (
	"fmt"

	"github.com/doruo/gameoflife/cell"
)

type Grid [][]cell.Cell

func NewGrid(len int) *Grid {
	m := make(Grid, len)
	for i := range len {
		m[i] = make([]cell.Cell, len)
	}
	return &m
}

func (m Grid) UpdateCellAdjs(i, j int) {

	adjs := make([]cell.Cell, 8)
	cell := m.GetCell(i, j)

	// For each possible adjacent position
	for _, position := range GetAdjacentsPos() {
		// Verify grid border case (to avoid Ci = -1)
		if m.IsValidPosition(i+position[0], j+position[1]) {
			// Appends each adjacent alive cell
			if adjCell := m.GetCell(i+position[0], j+position[1]); adjCell.IsAlive {
				adjs = append(adjs, *adjCell)
			}

		}
	}
	cell.SetAdjacent(adjs)
}

func (m Grid) UpdateCells() {
	for i := range m {
		for j := range m[i] {
			// For each alive cell
			c := m.GetCell(i, j)
			if c.IsAlive {
				m.UpdateCellAdjs(i, j)
				c.UpdateState()
				for _, cellAdj := range c.Adjacents {
					// Update cell and its adjacents
					cellAdj.UpdateState()
				}

			}
		}
	}
}

func (m Grid) Show() {
	for i := range m {
		for j := range m[i] {
			fmt.Print(m.GetCell(i, j).ToString(), " ")
		}
		fmt.Println(" ")
	}
}

func (m Grid) GetCell(i, j int) *cell.Cell {
	return &m[i][j]
}

// All possible adjacent cell position
func GetAdjacentsPos() [8][2]int {
	return [8][2]int{
		{-1, -1}, {-1, 0}, {-1, 1}, // upper level
		{0, -1}, {0, 1}, // left and right level
		{1, -1}, {1, 0}, {1, 1}} // bottom level
}

func (m Grid) IsValidPosition(i, j int) bool {
	return i >= 0 && i < len(m) && j >= 0 && j < len(m[0])
}
