package key_value

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestPipe(t *testing.T) {
	a, b := Pipe("a|b")
	assert.String(t, "a", a)
	assert.String(t, "b", b)
	c, d := Pipe("c|d|e")
	assert.String(t, "c", c)
	assert.String(t, "d|e", d)
	f, g := Pipe("f")
	assert.String(t, "f", f)
	assert.String(t, "", g)
	h, i := Pipe("h|")
	assert.String(t, "h", h)
	assert.String(t, "", i)
}
