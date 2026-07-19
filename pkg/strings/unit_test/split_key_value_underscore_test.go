package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings/split/key_value"
	"testing"
)

func TestUnderscore(t *testing.T) {
	a, b := key_value.Underscore("a_b")
	assert.String(t, "a", a)
	assert.String(t, "b", b)
	c, d := key_value.Underscore("c_d_e")
	assert.String(t, "c", c)
	assert.String(t, "d_e", d)
	f, g := key_value.Underscore("f")
	assert.String(t, "f", f)
	assert.String(t, "", g)
	h, i := key_value.Underscore("h_")
	assert.String(t, "h", h)
	assert.String(t, "", i)
}
