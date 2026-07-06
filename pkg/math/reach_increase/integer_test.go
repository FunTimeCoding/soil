package reach_increase

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestInteger(t *testing.T) {
	// Short by 1
	assertInteger(t, 48, 49, 50, false)
	// Reached exactly
	assertInteger(t, 49, 50, 50, true)
	// Exceed by 1
	assertInteger(t, 49, 51, 50, true)
	// Dip below
	assertInteger(t, 51, 49, 50, false)
}

func assertInteger(
	t *testing.T,
	past int,
	now int,
	threshold int,
	expect bool,
) {
	t.Helper()
	assert.Boolean(t, expect, Integer(past, now, threshold))
}
