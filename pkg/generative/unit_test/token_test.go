package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/generative/token"
	"testing"
)

func TestToken(t *testing.T) {
	assert.Integer(t, 3, token.Count("Hello world!"))
	assert.Integer(t, 2, token.CountSlice([]string{"Hello", "world"}))
}
