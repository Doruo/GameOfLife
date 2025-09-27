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

// --------------------------------------------

// Updates all cells, return all remaining alive cell positions
func (newGrid *Grid) UpdateCells(oldGrid *Grid) [][]int {

	for i := range *oldGrid {
		for j := range (*oldGrid)[i] {
			// Update cell and its adjacents
			newGrid.UpdateCell(i, j, oldGrid)
		}
	}

	return newGrid.GetAlivesPos()
}

// --------------------------------------------

func (newGrid *Grid) UpdateCell(i, j int, oldGrid *Grid) {

	oldGrid.UpdateCellAdjs(i, j)
	cOld := oldGrid.GetCell(i, j)
	cOld.UpdateState()
	newGrid.SetCell(i, j, *cOld)
}

// --------------------------------------------

func (oldGrid *Grid) UpdateCellAdjs(i, j int) {
	adjs := make([]cell.Cell, 0, 8)

	for _, position := range GetAdjacentsPos() {
		if oldGrid.IsValidPosition(i+position[0], j+position[1]) {
			if adjCell := oldGrid.GetCell(i+position[0], j+position[1]); adjCell.IsAlive() {
				adjs = append(adjs, *adjCell)
			}
		}
	}
	oldGrid.GetCell(i, j).SetAdjacent(adjs)
}

// --------------------------------------------

func (m Grid) Show() {
	for i := range m {
		for j := range m[i] {
			showCell(m.GetCell(i, j))
		}
		fmt.Println(" ")
	}
}

func showCell(cell *cell.Cell) {
	fmt.Print(cell.ToString(), " ")
}

// --------------------------------------------

func (m Grid) GetAlivesPos() [][]int {
	alives := make([][]int, len(m))
	for i := range m {
		for j := range m[i] {
			// For each cell
			c := m.GetCell(i, j)
			if c.IsAlive() {
				alives = append(alives, []int{i, j})
			}
		}
	}
	return alives
}

func (m Grid) GetCell(i, j int) *cell.Cell {
	return &m[i][j]
}

func (m Grid) SetCell(i, j int, c cell.Cell) {
	m[i][j] = c
}

func (m Grid) IsValidPosition(i, j int) bool {
	return i >= 0 && i < len(m) && j >= 0 && j < len(m[0])
}

// All possible adjacent cell position
func GetAdjacentsPos() [8][2]int {
	return [8][2]int{
		{-1, -1}, {-1, 0}, {-1, 1}, // upper level
		{0, -1}, {0, 1}, // left and right level
		{1, -1}, {1, 0}, {1, 1}} // bottom level
}
