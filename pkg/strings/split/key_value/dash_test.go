package key_value

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestDash(t *testing.T) {
	a, b := Dash("a-b")
	assert.String(t, "a", a)
	assert.String(t, "b", b)
	c, d := Dash("c-d-e")
	assert.String(t, "c", c)
	assert.String(t, "d-e", d)
	f, g := Dash("f")
	assert.String(t, "f", f)
	assert.String(t, "", g)
	h, i := Dash("h-")
	assert.String(t, "h", h)
	assert.String(t, "", i)
}
