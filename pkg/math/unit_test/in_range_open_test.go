package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/math/in_range"
	"github.com/funtimecoding/soil/pkg/math/ranges"
	"testing"
)

func TestOpen(t *testing.T) {
	zeroToOne := ranges.Range{L: 0, R: 1}
	assertOpen(t, 0, zeroToOne, false)
	assertOpen(t, 0.01, zeroToOne, true)
	assertOpen(t, 0.99, zeroToOne, true)
	assertOpen(t, 1, zeroToOne, false)
	assertOpen(t, 1.01, zeroToOne, false)
}

func assertOpen(
	t *testing.T,
	value float64,
	r ranges.Range,
	expect bool,
) {
	t.Helper()
	assert.Boolean(t, in_range.Open(value, r), expect)
}
