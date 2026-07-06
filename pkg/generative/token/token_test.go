package token

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestToken(t *testing.T) {
	assert.Integer(t, 3, Count("Hello world!"))
	assert.Integer(t, 2, CountSlice([]string{"Hello", "world"}))
}
