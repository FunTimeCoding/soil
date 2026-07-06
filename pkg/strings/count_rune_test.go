package strings

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestCountRune(t *testing.T) {
	assert.Integer(t, 0, CountRune("a", 'b'))
	assert.Integer(t, 1, CountRune("a", 'a'))
	assert.Integer(t, 2, CountRune("aa", 'a'))
}
