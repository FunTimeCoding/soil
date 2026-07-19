package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/math/slice_sum"
	"testing"
)

func TestSliceSumFloat(t *testing.T) {
	assert.Float(t, 3, slice_sum.Float([]float64{1, 2, 3}, 2))
}
