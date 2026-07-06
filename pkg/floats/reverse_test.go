package floats

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestReverse(t *testing.T) {
	reversed := []float64{1, 2, 3}
	Reverse(reversed)
	assert.Any(t, []float64{3, 2, 1}, reversed)
}
