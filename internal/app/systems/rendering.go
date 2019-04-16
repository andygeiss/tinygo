package systems

import (
	"github.com/andygeiss/tinygo/internal/app/views/layout"
	"github.com/andygeiss/tinygo/internal/pkg/ecs"
	"github.com/andygeiss/tinygo/internal/pkg/web/render"
	"github.com/andygeiss/tinygo/internal/pkg/web/style"
	"syscall/js"
)

// Rendering ...
type Rendering struct {
	width  int
	height int
}

// NewRendering ...
func NewRendering() ecs.System {
	return &Rendering{}
}

// Process ...
func (s *Rendering) Process(entityManager *ecs.EntityManager) {
	app := js.Global().Get("document").Call("getElementById", "app")
	// Render the content with a HTML5 Renderer.
	html := render.HTML5(entityManager)
	app.Set("innerHTML", html)
	// Layout (move, resize, ...)
	s.layoutAndStyle(entityManager)
}

// Setup ...
func (s *Rendering) Setup() {}

// Teardown ...
func (s *Rendering) Teardown() {}

func (s *Rendering) layoutAndStyle(entityManager *ecs.EntityManager) {
	// Get the current window size
	s.height = js.Global().Get("window").Get("innerHeight").Int()
	s.width = js.Global().Get("window").Get("innerWidth").Int()
	// Layout and style each element
	for _, entity := range entityManager.FilterBy("position", "size") {
		s.layoutElement(entity)
		ele := js.Global().Get("document").Call("getElementById", entity.Id)
		// Style the content with a CSS3 Styler.
		styles := style.CSS3(entity)
		ele.Set("style", styles)
	}
}

func (s *Rendering) layoutElement(entity *ecs.Entity) {
	if s.width > s.height {
		layout.LandscapeElement(s.width, s.height, entity)
	} else {
		layout.PortraitElement(s.width, s.height, entity)
	}
}
