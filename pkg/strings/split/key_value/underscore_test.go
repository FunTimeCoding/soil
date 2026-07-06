package key_value

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestUnderscore(t *testing.T) {
	a, b := Underscore("a_b")
	assert.String(t, "a", a)
	assert.String(t, "b", b)
	c, d := Underscore("c_d_e")
	assert.String(t, "c", c)
	assert.String(t, "d_e", d)
	f, g := Underscore("f")
	assert.String(t, "f", f)
	assert.String(t, "", g)
	h, i := Underscore("h_")
	assert.String(t, "h", h)
	assert.String(t, "", i)
}
