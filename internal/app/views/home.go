package views

import (
	"github.com/andygeiss/ecs"
	"github.com/andygeiss/tinygo/internal/app"
	"github.com/andygeiss/tinygo/internal/app/systems"
	"github.com/andygeiss/ui/entities"
)

type homeView struct{}

// NewHomeView ...
func NewHomeView() app.View {
	return &homeView{}
}

// Show ...
func (s *homeView) Show() {
	// Add entities and systems.
	entityManager := ecs.NewEntityManager()
	entityManager.Add(
		entities.NewTinyGoElement("content", "", "Hello Gophers!", "white", "blue", "homeSayHello();"),
	)
	systemManager := ecs.NewSystemManager()
	systemManager.Add(
		systems.NewRendering(),
	)
	// Finally setup, run and teardown the engine.
	engine := ecs.NewEngine(entityManager, systemManager)
	engine.Setup()
	engine.Run()
	engine.Teardown()
}
