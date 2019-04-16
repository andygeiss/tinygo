package render

import (
	"fmt"
	"github.com/andygeiss/tinygo/internal/pkg/ecs"
	"github.com/andygeiss/tinygo/internal/pkg/web/components"
)

// HTML5 ...
func HTML5(entityManager *ecs.EntityManager) (html string) {
	for _, entity := range entityManager.FilterBy("tag") {
		class := entity.Get("class")
		input := entity.Get("input")
		tag := entity.Get("tag").(*components.Tag)
		text := entity.Get("text")
		innerHTML := ""
		if text != nil {
			innerHTML = text.(*components.Text).Value
		}
		attributes := ""
		if class != nil {
			cl := class.(*components.Class)
			if cl.List != "" {
				attributes += fmt.Sprintf(` class="%s"`, cl.List)
			}
		}
		// HTML5 need a way to add Go functions to each HTML5 element without using "addEventListener".
		if input != nil {
			in := input.(*components.Input)
			if in.OnClick != "" {
				attributes += fmt.Sprintf(` onclick="%s"`, in.OnClick)
			}
		}
		html += fmt.Sprintf(`<%s id="%s"%s>%s</%s>`, tag.Value, entity.Id, attributes, innerHTML, tag.Value)
	}
	return html
}
