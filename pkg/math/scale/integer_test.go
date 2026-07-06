package scale

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestInteger(t *testing.T) {
	assertInteger(t, 0, 100, 0, 0)
	assertInteger(t, 0, 100, 0.5, 50)
	assertInteger(t, 0, 100, 1, 100)
	assertInteger(t, 40, 20, 0.5, 30)
	assertInteger(t, -10, -20, 0.5, -15)
	assertInteger(t, -30, -10, 0.5, -20)
}

func assertInteger(
	t *testing.T,
	from int,
	to int,
	factor float64,
	expect int,
) {
	t.Helper()
	assert.Integer(t, expect, Integer(from, to, factor))
}
