package key_value

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestColon(t *testing.T) {
	a, b := Colon("a:b")
	assert.String(t, "a", a)
	assert.String(t, "b", b)
	c, d := Colon("c:d:e")
	assert.String(t, "c", c)
	assert.String(t, "d:e", d)
	f, g := Colon("f")
	assert.String(t, "f", f)
	assert.String(t, "", g)
	h, i := Colon("h:")
	assert.String(t, "h", h)
	assert.String(t, "", i)
}
