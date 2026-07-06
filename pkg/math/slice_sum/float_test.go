package slice_sum

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestFloat(t *testing.T) {
	assert.Float(t, 3, Float([]float64{1, 2, 3}, 2))
}
