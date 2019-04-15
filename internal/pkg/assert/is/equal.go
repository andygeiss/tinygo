package is

import (
	"github.com/andygeiss/tinygo/internal/pkg/assert"
	"reflect"
)

// equalMatcher with its value will not be exported.
type equalMatcher struct {
	val interface{}
}

// Matches checks if values are equal.
func (m *equalMatcher) Matches(val interface{}) bool {
	return reflect.DeepEqual(m.val, val)
}

// Name returns the unique name of the matcher.
func (m *equalMatcher) Name() string {
	return "equalMatcher"
}

// Value ...
func (m *equalMatcher) Value() interface{} {
	return m.val
}

// Equal initializes the Equal-Matcher.
func Equal(val interface{}) assert.Matcher {
	return &equalMatcher{val: val}
}
