package key_value

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestDot(t *testing.T) {
	a, b := Dot("a.b")
	assert.String(t, "a", a)
	assert.String(t, "b", b)
	c, d := Dot("c.d.e")
	assert.String(t, "c", c)
	assert.String(t, "d.e", d)
	f, g := Dot("f")
	assert.String(t, "f", f)
	assert.String(t, "", g)
	h, i := Dot("h.")
	assert.String(t, "h", h)
	assert.String(t, "", i)
}
