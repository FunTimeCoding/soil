package key_value

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestAt(t *testing.T) {
	a, b := At("a@b")
	assert.String(t, "a", a)
	assert.String(t, "b", b)
	c, d := At("c@d@e")
	assert.String(t, "c", c)
	assert.String(t, "d@e", d)
	f, g := At("f")
	assert.String(t, "f", f)
	assert.String(t, "", g)
	h, i := At("h@")
	assert.String(t, "h", h)
	assert.String(t, "", i)
}
