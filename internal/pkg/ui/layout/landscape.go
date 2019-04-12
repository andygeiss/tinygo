package layout

import (
	"github.com/andygeiss/tinygo/internal/pkg/ecs"
	"github.com/andygeiss/tinygo/internal/pkg/ui/components"
)

// LandscapeElement ...
func LandscapeElement(maxWidth, maxHeight int, entity *ecs.Entity) {
	pos := entity.Get("position").(*components.Position)
	size := entity.Get("size").(*components.Size)
	size.Height = maxHeight * 10 / 100
	size.Width = maxWidth * 90 / 100
	pos.X = maxWidth * 30 / 100
	pos.Y = maxHeight * 30 / 100
}
