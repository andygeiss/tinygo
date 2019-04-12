package render_test

import (
	"github.com/andygeiss/tinygo/internal/pkg/assert"
	"github.com/andygeiss/tinygo/internal/pkg/assert/is"
	"github.com/andygeiss/tinygo/internal/pkg/ecs"
	"github.com/andygeiss/tinygo/internal/pkg/ui/entities"
	"github.com/andygeiss/tinygo/internal/pkg/ui/render"
	"testing"
)

func TestHTML_Should_Render_Element_Without_Text(t *testing.T) {
	em := ecs.NewEntityManager()
	em.Add(
		entities.NewElement("foo", "", "", "", "", "alert(1)"),
	)
	html := render.HTML5(em)
	assert.That(t, html, is.Equal(`<div id="foo" onclick="alert(1)"></div>`))
}

func TestHTML_Should_Render_Element_With_Text(t *testing.T) {
	em := ecs.NewEntityManager()
	em.Add(
		entities.NewElement("foo", "", "bar", "", "", "alert(1)"),
	)
	html := render.HTML5(em)
	assert.That(t, html, is.Equal(`<div id="foo" onclick="alert(1)">bar</div>`))
}

func TestHTML_Should_Render_Element_With_Class(t *testing.T) {
	em := ecs.NewEntityManager()
	em.Add(
		entities.NewElement("foo", "bar", "", "", "", "alert(1)"),
	)
	html := render.HTML5(em)
	assert.That(t, html, is.Equal(`<div id="foo" class="bar" onclick="alert(1)"></div>`))
}
