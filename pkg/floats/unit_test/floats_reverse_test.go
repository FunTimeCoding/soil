package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/floats"
	"testing"
)

func TestReverse(t *testing.T) {
	reversed := []float64{1, 2, 3}
	floats.Reverse(reversed)
	assert.Any(t, []float64{3, 2, 1}, reversed)
}
