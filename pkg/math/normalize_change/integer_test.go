package normalize_change

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestInteger(t *testing.T) {
	assertInteger(t, 0, 1, 0, 100, 1)
	assertInteger(t, 1, -2, 0, 100, -1)
	assertInteger(t, 1, -3, 0, 100, -1)
	assertInteger(t, 100, -100, 0, 100, -100)
	assertInteger(t, 100, -150, 0, 100, -100)
	assertInteger(t, 95, 20, 0, 100, 5)
}

func TestIntegerNoMaximum(t *testing.T) {
	assertInteger(t, 0, 1, 0, 0, 1)
	assertInteger(t, 1, -2, 0, 0, -1)
	assertInteger(t, 1, -3, 0, 0, -1)
	assertInteger(t, 100, -100, 0, 0, -100)
	assertInteger(t, 100, -150, 0, 0, -100)
}

func assertInteger(
	t *testing.T,
	now int,
	change int,
	minimum int,
	maximum int,
	expect int,
) {
	t.Helper()
	assert.Integer(t, expect, Integer(now, change, minimum, maximum))
}
