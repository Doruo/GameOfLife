package cell

import "github.com/doruo/gameoflife/game/color"

type Cell struct {
	IsAlive   bool
	Adjacents []Cell
}

// Creates a new cell with alive adjacents in memory.
func NewCell() *Cell {
	return &Cell{
		IsAlive:   false,
		Adjacents: make([]Cell, 8),
	}
}

// Update cell state based on Conway's rules
func (c *Cell) UpdateState() {
	c.IsAlive = c.GetUpdatedState()
}

// Returns updated cell state based on Conway's rules.
func (c *Cell) GetUpdatedState() bool {
	return (len(c.Adjacents) == 3) || (len(c.Adjacents) == 2 && c.IsAlive)
}

func (c *Cell) IsAdjacent(c2 Cell) bool {
	for _, cell := range c.Adjacents {
		if &cell == &c2 {
			return true
		}
	}
	return false
}

func (c *Cell) SetAdjacent(adjs []Cell) {
	c.Adjacents = adjs
}

func (c *Cell) ToString() string {
	if c.IsAlive {
		return (color.Green + "O" + color.Reset)
	}
	return (color.Red + "~" + color.Reset)
}
