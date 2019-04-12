package style

import (
	"github.com/andygeiss/ecs"
	"github.com/andygeiss/tinygo/internal/pkg/ui/components"
	"sort"
	"strconv"
	"strings"
)

// CSS3 ...
func CSS3(entity *ecs.Entity) string {
	background := entity.Get("background")
	pos := entity.Get("position").(*components.Position)
	size := entity.Get("size").(*components.Size)
	text := entity.Get("text")
	width := strconv.Itoa(size.Width)
	height := strconv.Itoa(size.Height)
	mapping := map[string]string{
		"top":      strconv.Itoa(pos.Y) + "px",
		"left":     strconv.Itoa(pos.X) + "px",
		"position": "absolute",
		"height":   height + "px",
		"width":    width + "px",
	}
	if background != nil {
		bg := background.(*components.Background)
		if bg.Color != "" {
			mapping["background-color"] = bg.Color
		}
	}
	if text != nil {
		txt := text.(*components.Text)
		if txt.Value != "" {
			if txt.Color != "" {
				mapping["color"] = txt.Color
			}
		}
	}
	var index []string
	var styles []string
	for k := range mapping {
		index = append(index, k)
	}
	sort.Strings(index)
	for _, k := range index {
		styles = append(styles, k+": "+mapping[k])
	}
	style := strings.Join(styles, ";")
	return style
}
