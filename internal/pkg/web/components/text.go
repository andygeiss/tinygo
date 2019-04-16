package components

// Text ...
type Text struct {
	Color string
	Value string
}

// Name ...
func (c *Text) Name() string {
	return "text"
}
