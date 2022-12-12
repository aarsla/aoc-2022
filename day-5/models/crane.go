package models

import (
	"sort"
)

type Crane struct {
	Ship map[int][]string
}

func NewCrane() Crane {
	initialState := map[int][]string{
		1: []string{"D", "H", "N", "Q", "T", "W", "V", "B"},
		2: []string{"D", "W", "B"},
		3: []string{"T", "S", "Q", "W", "J", "C"},
		4: []string{"F", "J", "R", "N", "Z", "T", "P"},
		5: []string{"G", "P", "V", "J", "M", "S", "T"},
		6: []string{"B", "W", "F", "T", "N"},
		7: []string{"B", "L", "D", "Q", "F", "H", "V", "N"},
		8: []string{"H", "P", "F", "R"},
		9: []string{"Z", "S", "M", "B", "L", "N", "P", "H"},
	}

	return Crane{
		Ship: initialState,
	}
}

func (c *Crane) Move(ins Instruction) {
	col := ins.From
	newCol := ins.To

	for i := 1; i <= ins.Pieces; i++ {
		// Read a crate
		crate := c.Ship[col][len(c.Ship[col])-1]
		// Remove crate from col
		c.Ship[col] = c.Ship[col][:len(c.Ship[col])-1]
		// Add crate to new col
		c.Ship[newCol] = append(c.Ship[newCol], crate)
	}
}

func (c *Crane) Move9001(ins Instruction) {
	col := ins.From
	newCol := ins.To

	var crates []string
	for i := 1; i <= ins.Pieces; i++ {
		// Read a crate
		crate := c.Ship[col][len(c.Ship[col])-1]
		// Store it for moving
		crates = append(crates, crate)
		// Remove crate from col
		c.Ship[col] = c.Ship[col][:len(c.Ship[col])-1]
	}

	// Move pieces to a new col
	for i := len(crates) - 1; i >= 0; i-- {
		c.Ship[newCol] = append(c.Ship[newCol], crates[i])
	}
}

func (c *Crane) TopOfStack() string {
	var topOfStack string

	keys := make([]int, 0, len(c.Ship))
	for k := range c.Ship {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, k := range keys {
		topOfStack += c.Ship[k][len(c.Ship[k])-1]
	}

	return topOfStack
}
