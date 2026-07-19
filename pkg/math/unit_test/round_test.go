package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/math/round"
	"testing"
)

func TestRound(t *testing.T) {
	assert.Float(t, 1.0, round.Round(1.01, 1))
}
