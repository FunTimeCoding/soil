package key_value

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestSlash(t *testing.T) {
	a, b := Slash("a/b")
	assert.String(t, "a", a)
	assert.String(t, "b", b)
	c, d := Slash("c/d/e")
	assert.String(t, "c", c)
	assert.String(t, "d/e", d)
	f, g := Slash("f")
	assert.String(t, "f", f)
	assert.String(t, "", g)
	h, i := Slash("h/")
	assert.String(t, "h", h)
	assert.String(t, "", i)
}
