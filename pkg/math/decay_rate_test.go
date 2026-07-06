package math

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestDecayRate(t *testing.T) {
	assert.Round(t, 0.07, DecayRate(10), 2)
}
