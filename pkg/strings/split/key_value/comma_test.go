package key_value

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestComma(t *testing.T) {
	a, b := Comma("a,b")
	assert.String(t, "a", a)
	assert.String(t, "b", b)
	c, d := Comma("c,d,e")
	assert.String(t, "c", c)
	assert.String(t, "d,e", d)
	f, g := Comma("f")
	assert.String(t, "f", f)
	assert.String(t, "", g)
	h, i := Comma("h,")
	assert.String(t, "h", h)
	assert.String(t, "", i)
}
