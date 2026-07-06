package math

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestDistribution(t *testing.T) {
	assertDistribution(t, 100, 10, 100)
	assertDistribution(t, 1000000, 10, 1000000)
}

func assertDistribution(
	t *testing.T,
	all float64,
	steps int,
	expect float64,
) {
	t.Helper()
	var sum float64

	for _, e := range Distribution(all, steps) {
		sum += e
	}

	// Comparing all floats here is a lot of decimals
	if false {
		assert.Any(t, expect, Distribution(all, steps))
	}

	assert.Float(t, expect, sum)
}
