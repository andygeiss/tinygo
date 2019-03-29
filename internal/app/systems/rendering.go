package systems

import (
	"github.com/andygeiss/ecs"
	"github.com/andygeiss/ui/layout"
	"github.com/andygeiss/ui/render"
	"github.com/andygeiss/ui/style"
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
	// Render
	app := js.Global().Get("document").Call("getElementById", "app")
	html := render.TinyGo(entityManager)
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
		styles := style.Default(entity)
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
