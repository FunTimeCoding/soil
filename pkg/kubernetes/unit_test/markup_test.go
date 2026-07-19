package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/kubernetes/markup"
	"testing"
)

func TestRender(t *testing.T) {
	assert.Any(t, "null\n", markup.Render(nil))
}
