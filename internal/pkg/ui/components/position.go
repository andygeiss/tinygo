package components

// Position ...
type Position struct {
	X int
	Y int
}

// Name ...
func (c *Position) Name() string {
	return "position"
}
