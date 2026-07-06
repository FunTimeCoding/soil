package key_value

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestSemicolon(t *testing.T) {
	a, b := Semicolon("a;b")
	assert.String(t, "a", a)
	assert.String(t, "b", b)
	c, d := Semicolon("c;d;e")
	assert.String(t, "c", c)
	assert.String(t, "d;e", d)
	f, g := Semicolon("f")
	assert.String(t, "f", f)
	assert.String(t, "", g)
	h, i := Semicolon("h;")
	assert.String(t, "h", h)
	assert.String(t, "", i)
}
