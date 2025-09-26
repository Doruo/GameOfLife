package cell

type Cell struct {
	isAlive bool
	adj     []Cell
}

func NewCell() *Cell {
	return &Cell{
		isAlive: false,
		adj:     make([]Cell, 8),
	}
}

func (c *Cell) UpdateState() {
	c.SetAlive(c.GetUpdatedState())
}

// Returns updated cell state based on Conway's rules.
func (c *Cell) GetUpdatedState() bool {
	return (len(c.adj) == 3) || (len(c.adj) == 2 && c.isAlive)
}

// Get current state (false => dead, true => alive).
func (c *Cell) GetState() bool {
	return c.isAlive
}

func (c *Cell) SetAlive(state bool) {
	c.isAlive = state
}
