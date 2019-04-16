package layout_test

import (
	"github.com/andygeiss/tinygo/internal/app/views"
	"github.com/andygeiss/tinygo/internal/pkg/assert"
	"github.com/andygeiss/tinygo/internal/pkg/assert/is"
	"github.com/andygeiss/tinygo/internal/pkg/web/components"
	"testing"
)

func TestLandscapeElement(t *testing.T) {
	maxWidth := 20
	maxHeight := 10
	eleHeight := maxHeight * 10 / 100
	eleWidth := maxWidth * 90 / 100
	posX := maxWidth * 30 / 100
	posY := maxHeight * 30 / 100
	e := views.NewElement("e", "", "foo", "", "", "")
	views.LandscapeElement(maxWidth, maxHeight, e)
	pos := e.Get("position").(*components.Position)
	size := e.Get("size").(*components.Size)
	assert.That(t, pos.X, is.Equal(posX))
	assert.That(t, pos.Y, is.Equal(posY))
	assert.That(t, size.Height, is.Equal(eleHeight))
	assert.That(t, size.Width, is.Equal(eleWidth))
}
