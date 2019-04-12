package style_test

import (
	"github.com/andygeiss/tinygo/internal/pkg/assert"
	"github.com/andygeiss/tinygo/internal/pkg/assert/is"
	"github.com/andygeiss/tinygo/internal/pkg/ui/entities"
	"github.com/andygeiss/tinygo/internal/pkg/ui/style"
	"testing"
)

func TestDefault_Should_Return_Styled_Element(t *testing.T) {
	ele := entities.NewElement("foo", "", "", "red", "black", "")
	styles := style.CSS3(ele)
	assert.That(t, styles, is.Equal(`background-color: black;height: 0px;left: 0px;position: absolute;top: 0px;width: 0px`))
}

func TestDefault_Should_Return_Styled_Element_Text(t *testing.T) {
	ele := entities.NewElement("foo", "", "bar", "red", "black", "")
	styles := style.CSS3(ele)
	assert.That(t, styles, is.Equal(`background-color: black;color: red;height: 0px;left: 0px;position: absolute;top: 0px;width: 0px`))
}
