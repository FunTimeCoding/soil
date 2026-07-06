package key_value

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestSpace(t *testing.T) {
	a, b := Space("a b")
	assert.String(t, "a", a)
	assert.String(t, "b", b)
	c, d := Space("c d e")
	assert.String(t, "c", c)
	assert.String(t, "d e", d)
	f, g := Space("f")
	assert.String(t, "f", f)
	assert.String(t, "", g)
	h, i := Space("h ")
	assert.String(t, "h", h)
	assert.String(t, "", i)
}
