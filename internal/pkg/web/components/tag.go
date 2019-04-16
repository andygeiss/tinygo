package components

// Tag ...
type Tag struct {
	Value string
}

// Name ...
func (c *Tag) Name() string {
	return "tag"
}
