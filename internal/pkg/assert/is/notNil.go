package is

import (
	"github.com/andygeiss/tinygo/internal/pkg/assert"
	"reflect"
)

// notNilMatcher with its value will not be exported.
type notNilMatcher struct{}

// Matches checks if values are notNil.
func (m *notNilMatcher) Matches(val interface{}) bool {
	return !reflect.DeepEqual(nil, val)
}

// Name returns the unique name of the matcher.
func (m *notNilMatcher) Name() string {
	return "notNilMatcher"
}

// Value ...
func (m *notNilMatcher) Value() interface{} {
	return "not nil"
}

// NotNil initializes the notNil-Matcher.
func NotNil() assert.Matcher {
	return &notNilMatcher{}
}
