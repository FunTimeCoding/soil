package scale

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestFloat(t *testing.T) {
	assertFloat(t, 0, 100, 0, 0)
	assertFloat(t, 0, 100, 0.5, 50)
	assertFloat(t, 0, 100, 1, 100)
	assertFloat(t, 40, 20, 0.5, 30)
	assertFloat(t, 1, 1, 0.5, 1)
	assertFloat(t, -1, 1, 0.5, 0)
	assertFloat(t, 2.0/3.0, 1, 0, 0.7)
	assertFloat(t, 2.0/3.0, 1, 1, 1)
}

func assertFloat(
	t *testing.T,
	from float64,
	to float64,
	factor float64,
	expect float64,
) {
	t.Helper()
	assert.Round(t, expect, Float(from, to, factor), 1)
}
