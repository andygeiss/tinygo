package layout_test

import (
	"github.com/andygeiss/tinygo/internal/app/views"
	"github.com/andygeiss/tinygo/internal/pkg/assert"
	"github.com/andygeiss/tinygo/internal/pkg/assert/is"
	"github.com/andygeiss/tinygo/internal/pkg/web/components"
	"testing"
)

func TestPortraitElement(t *testing.T) {
	maxWidth := 10
	maxHeight := 20
	eleHeight := maxHeight * 20 / 100
	eleWidth := maxWidth * 60 / 100
	posX := maxWidth * 20 / 100
	posY := maxHeight * 40 / 100
	e := views.NewElement("e", "", "foo", "", "", "")
	views.PortraitElement(maxWidth, maxHeight, e)
	pos := e.Get("position").(*components.Position)
	size := e.Get("size").(*components.Size)
	assert.That(t, pos.X, is.Equal(posX))
	assert.That(t, pos.Y, is.Equal(posY))
	assert.That(t, size.Height, is.Equal(eleHeight))
	assert.That(t, size.Width, is.Equal(eleWidth))
}
