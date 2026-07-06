package math

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestLogarithm(t *testing.T) {
	assert.Float(t, 0, Logarithm(1))
}
