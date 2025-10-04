package game

import (
	"fmt"
)

type Grid [][]Cell

func NewGrid(len int) *Grid {
	m := make(Grid, len)
	for i := range len {
		m[i] = make([]Cell, len)
	}
	return &m
}

// --------------------------------------------

// Updates all cells, returns all remaining alives
func (newGrid *Grid) UpdateCells(oldGrid *Grid) [][]int {

	alives := make([][]int, len(*oldGrid))
	for i := range *oldGrid {
		for j := range (*oldGrid)[i] {
			// Update cell and its adjacents
			newGrid.UpdateCell(i, j, oldGrid)
			if oldGrid.GetCell(i, j).IsAlive() {
				alives = append(alives, []int{i, j})
			}
		}
	}
	return alives
}

// Updates a given cell
func (newGrid *Grid) UpdateCell(i, j int, oldGrid *Grid) {
	oldGrid.UpdateCellAdjs(i, j)

	cOld := oldGrid.GetCell(i, j)
	cOld.UpdateState()

	newGrid.SetCell(i, j, *cOld)
}

func (oldGrid *Grid) UpdateCellAdjs(i, j int) {
	adjs := make([]Cell, 0, 8)

	for _, position := range GetAdjacentsPos() {
		if oldGrid.IsValidPosition(i+position[0], j+position[1]) {
			if adjCell := oldGrid.GetCell(i+position[0], j+position[1]); adjCell.IsAlive() {
				adjs = append(adjs, *adjCell)
			}
		}
	}
	oldGrid.GetCell(i, j).SetAdjacents(adjs)
}

// --------------------------------------------

func (m Grid) Display() {
	for i := range m {
		for j := range m[i] {
			displayCell(m.GetCell(i, j))
		}
		fmt.Println(" ")
	}
}

func displayCell(cell *Cell) {
	fmt.Print(cell.ToString(), " ")
}

// --------------------------------------------

func (m Grid) GetCell(i, j int) *Cell {
	return &m[i][j]
}

func (m Grid) SetCell(i, j int, c Cell) {
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
