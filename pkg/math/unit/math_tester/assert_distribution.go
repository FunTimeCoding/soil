package math_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/math"
	"testing"
)

func AssertDistribution(
	t *testing.T,
	expected float64,
	all float64,
	steps int,
) {
	t.Helper()
	var sum float64

	for _, e := range math.Distribution(all, steps) {
		sum += e
	}

	assert.Float(t, expected, sum)
}
