package is

import (
	"github.com/andygeiss/tinygo/internal/pkg/assert"
	"reflect"
)

type notEqualMatcher struct {
	val interface{}
}

// Matches ...
func (m *notEqualMatcher) Matches(val interface{}) bool {
	return !reflect.DeepEqual(m.val, val)
}

// Name returns the unique name of the matcher.
func (m *notEqualMatcher) Name() string {
	return "notEqualMatcher"
}

// Value ...
func (m *notEqualMatcher) Value() interface{} {
	return m.val
}

// NotEqual ...
func NotEqual(val interface{}) assert.Matcher {
	return &notEqualMatcher{val: val}
}
