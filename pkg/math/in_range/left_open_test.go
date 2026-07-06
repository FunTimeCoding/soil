package in_range

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/math/ranges"
	"testing"
)

func TestLeftOpen(t *testing.T) {
	zeroToOne := ranges.Range{L: 0, R: 1}
	assertLeftOpen(t, 0, zeroToOne, false)
	assertLeftOpen(t, 0.01, zeroToOne, true)
	assertLeftOpen(t, 0.99, zeroToOne, true)
	assertLeftOpen(t, 1, zeroToOne, true)
	assertLeftOpen(t, 1.01, zeroToOne, false)
}

func assertLeftOpen(
	t *testing.T,
	value float64,
	r ranges.Range,
	expect bool,
) {
	t.Helper()
	assert.Boolean(t, LeftOpen(value, r), expect)
}
