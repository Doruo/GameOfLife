package cell

type Cell struct {
	isAlive   bool
	adjacents []Cell
}

// Creates a new cell with 8 adjacents in memory, all dead by default.
func NewCell() *Cell {
	return &Cell{
		isAlive:   false,
		adjacents: make([]Cell, 8),
	}
}

// Update cell state based on Conway's rules
func (c *Cell) UpdateState() {
	c.SetAlive(c.GetUpdatedState())
}

// Returns updated cell state based on Conway's rules.
func (c *Cell) GetUpdatedState() bool {
	return (len(c.adjacents) == 3) || (len(c.adjacents) == 2 && c.isAlive)
}

// Get current state (false => dead, true => alive).
func (c *Cell) IsAlive() bool {
	return c.isAlive
}

// Get current state (false => dead, true => alive).
func (c *Cell) GetAdjacents() []Cell {
	return c.adjacents
}

func (c *Cell) IsAdjacent(c2 Cell) bool {
	for _, cell := range c.adjacents {
		if &cell == &c2 {
			return true
		}
	}
	return false
}

func (c *Cell) SetAlive(state bool) {
	c.isAlive = state
}

func (c *Cell) SetAdjacents(adjs []Cell) {
	c.adjacents = adjs
}
