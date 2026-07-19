package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings"
	"testing"
)

func TestCountRune(t *testing.T) {
	assert.Integer(t, 0, strings.CountRune("a", 'b'))
	assert.Integer(t, 1, strings.CountRune("a", 'a'))
	assert.Integer(t, 2, strings.CountRune("aa", 'a'))
}
