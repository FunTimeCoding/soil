package key_value

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestEquals(t *testing.T) {
	a, b := Equals("a=b")
	assert.String(t, "a", a)
	assert.String(t, "b", b)
	c, d := Equals("c=d=e")
	assert.String(t, "c", c)
	assert.String(t, "d=e", d)
	f, g := Equals("f")
	assert.String(t, "f", f)
	assert.String(t, "", g)
	h, i := Equals("h=")
	assert.String(t, "h", h)
	assert.String(t, "", i)
}
