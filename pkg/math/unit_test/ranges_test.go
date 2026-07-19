package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/math/ranges"
	"testing"
)

func TestRange(t *testing.T) {
	assert.Any(t, &ranges.Range{L: 1, R: 2}, ranges.New(1, 2))
	assert.Any(t, &ranges.Range{L: -1, R: 3}, ranges.FromFactor(1, 2))
}
