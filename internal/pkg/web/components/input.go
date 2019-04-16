package components

// Input ...
type Input struct {
	Click   func()
	OnClick string
}

// Name ...
func (c *Input) Name() string {
	return "input"
}
