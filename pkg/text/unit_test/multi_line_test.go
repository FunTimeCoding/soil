package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/text/multi_line"
	"testing"
)

func TestMultiLine(t *testing.T) {
	l := multi_line.New()
	assert.True(t, l.Empty())
	l.Add("a")
	assert.False(t, l.Empty())
	l.Add("b")
	assert.String(t, "a\nb", l.Render())
}
