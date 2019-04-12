package entities

import (
	"github.com/andygeiss/ecs"
	"github.com/andygeiss/tinygo/internal/pkg/ui/components"
)

// NewElement ...
func NewElement(id, classList, text, textColor, backgroundColor, onClick string) (entity *ecs.Entity) {
	return &ecs.Entity{
		Id: id,
		Components: []ecs.Component{
			&components.Background{Color: backgroundColor},
			&components.Class{List: classList},
			&components.Input{OnClick: onClick},
			&components.Position{X: 0, Y: 0},
			&components.Size{Width: 0, Height: 0},
			&components.Tag{Value: "div"},
			&components.Text{Color: textColor, Value: text},
		},
	}
}
