package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/math"
	"testing"
)

func TestDecayRate(t *testing.T) {
	assert.Round(t, 0.07, math.DecayRate(10), 2)
}
