package key_value

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestTab(t *testing.T) {
	a, b := Tab("a\tb")
	assert.String(t, "a", a)
	assert.String(t, "b", b)
	c, d := Tab("c\td\te")
	assert.String(t, "c", c)
	assert.String(t, "d\te", d)
	f, g := Tab("f")
	assert.String(t, "f", f)
	assert.String(t, "", g)
	h, i := Tab("h\t")
	assert.String(t, "h", h)
	assert.String(t, "", i)
}
