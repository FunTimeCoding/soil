package multi_line

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestMultiLine(t *testing.T) {
	l := New()
	assert.True(t, l.Empty())
	l.Add("a")
	assert.False(t, l.Empty())
	l.Add("b")
	assert.String(t, "a\nb", l.Render())
}
