package reach_increase

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestFloat(t *testing.T) {
	// Short by 1
	assertFloat(t, 48, 49, 50, false)
	// Reached exactly
	assertFloat(t, 49, 50, 50, true)
	// Exceed by 1
	assertFloat(t, 49, 51, 50, true)
	// Reach decrease
	assertFloat(t, 51, 49, 50, false)
}

func assertFloat(
	t *testing.T,
	past float64,
	now float64,
	threshold float64,
	expect bool,
) {
	t.Helper()
	assert.Boolean(t, expect, Float(past, now, threshold))
}
