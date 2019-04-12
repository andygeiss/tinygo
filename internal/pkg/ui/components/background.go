package components

// Background ...
type Background struct {
	Color string
}

// Name ...
func (c *Background) Name() string {
	return "background"
}
