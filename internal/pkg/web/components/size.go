package components

// Size ...
type Size struct {
	Width  int
	Height int
}

// Name ...
func (c *Size) Name() string {
	return "size"
}
