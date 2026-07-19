package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/math"
	"testing"
)

func TestLogarithm(t *testing.T) {
	assert.Float(t, 0, math.Logarithm(1))
}
