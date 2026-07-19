package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings/split/key_value"
	"testing"
)

func TestSemicolon(t *testing.T) {
	a, b := key_value.Semicolon("a;b")
	assert.String(t, "a", a)
	assert.String(t, "b", b)
	c, d := key_value.Semicolon("c;d;e")
	assert.String(t, "c", c)
	assert.String(t, "d;e", d)
	f, g := key_value.Semicolon("f")
	assert.String(t, "f", f)
	assert.String(t, "", g)
	h, i := key_value.Semicolon("h;")
	assert.String(t, "h", h)
	assert.String(t, "", i)
}
