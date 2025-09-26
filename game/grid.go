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

func (m Grid) InitiateCellAdjs(i, j int) {

	adjs := make([]cell.Cell, 8)

	adjs[0] = m.GetCell(i+1, j-1)
	adjs[1] = m.GetCell(i-1, j)
	adjs[2] = m.GetCell(i+1, j+1)
	adjs[3] = m.GetCell(i, j-1)

	adjs[4] = m.GetCell(i, j+1)
	adjs[5] = m.GetCell(i-1, j-1)
	adjs[6] = m.GetCell(i+1, j)
	adjs[7] = m.GetCell(i-1, j+1)

	cell := m.GetCell(i, j)
	cell.SetAdjacents(adjs)
}

func (m Grid) UpdateCells() {
	for i := range m {
		for j := range m[i] {
			c1 := m.GetCell(i, j)
			if c1.IsAlive() {
				for _, c2 := range c1.GetAdjacents() {
					c2.UpdateState()
				}
			}
		}
	}
}

func (m Grid) Show() {
	for i := range m {
		for j := range m[i] {
			fmt.Print(m[i][j], " ")
		}
		fmt.Println(" ")
	}
}

func (m Grid) GetCell(i, j int) cell.Cell {
	return m[i][j]
}
