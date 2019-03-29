package services

import (
	"github.com/andygeiss/ecs"
	"github.com/andygeiss/tinygo/internal/app"
	"github.com/andygeiss/tinygo/internal/app/systems"
	"github.com/andygeiss/ui/entities"
	"syscall/js"
)

type homeService struct{}

// NewHomeService ...
func NewHomeService() app.Service {
	return &homeService{}
}

// Run ...
func (s *homeService) Run() {
	// Add entities and systems.
	entityManager := ecs.NewEntityManager()
	entityManager.Add(
		entities.NewTinyGoElement("content", "", "Hello Gophers!", "white", "blue", "homeServiceSend();"),
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

//go:export homeServiceSend
func homeServiceSend() {
	js.Global().Call("alert", "alert called from Go!")
}
