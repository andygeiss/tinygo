package is

import (
	"github.com/andygeiss/tinygo/internal/pkg/assert"
	"reflect"
)

// nilMatcher with its value will not be exported.
type nilMatcher struct{}

// Matches checks if values are nil.
func (m *nilMatcher) Matches(val interface{}) bool {
	return reflect.DeepEqual(nil, val)
}

// Name returns the unique name of the matcher.
func (m *nilMatcher) Name() string {
	return "nilMatcher"
}

// Value ...
func (m *nilMatcher) Value() interface{} {
	return nil
}

// Nil initializes the nil-Matcher.
func Nil() assert.Matcher {
	return &nilMatcher{}
}
