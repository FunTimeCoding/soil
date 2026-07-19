package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings/split/key_value"
	"testing"
)

func TestTab(t *testing.T) {
	a, b := key_value.Tab("a\tb")
	assert.String(t, "a", a)
	assert.String(t, "b", b)
	c, d := key_value.Tab("c\td\te")
	assert.String(t, "c", c)
	assert.String(t, "d\te", d)
	f, g := key_value.Tab("f")
	assert.String(t, "f", f)
	assert.String(t, "", g)
	h, i := key_value.Tab("h\t")
	assert.String(t, "h", h)
	assert.String(t, "", i)
}
