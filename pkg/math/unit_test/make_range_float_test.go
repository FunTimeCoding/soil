package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/math/make_range"
	"testing"
)

func TestMakeRangeFloat(t *testing.T) {
	assert.Any(t, []float64{0, 1}, make_range.Float(0, 1))
}
