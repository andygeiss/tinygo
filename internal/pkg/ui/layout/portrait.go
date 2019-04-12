package layout

import (
	"github.com/andygeiss/ecs"
	"github.com/andygeiss/tinygo/internal/pkg/ui/components"
)

// PortraitElement ...
func PortraitElement(maxWidth, maxHeight int, entity *ecs.Entity) {
	pos := entity.Get("position").(*components.Position)
	size := entity.Get("size").(*components.Size)
	size.Height = maxHeight * 20 / 100
	size.Width = maxWidth * 60 / 100
	pos.X = maxWidth * 20 / 100
	pos.Y = maxHeight * 40 / 100
}
