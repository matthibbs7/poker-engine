package chips

type Chips struct {
	Amount int
}

func NewChips(n int) Chips {
	return Chips{Amount: n}
}

func (c *Chips) Add(n int) {
	c.Amount += n
}

func (c *Chips) Remove(n int) {
	c.Amount -= n
}

func (c Chips) CanAfford(n int) bool {
	return c.Amount >= n
}
