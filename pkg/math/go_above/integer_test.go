package go_above

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestInteger(t *testing.T) {
	assertInteger(t, 52, 51, 50, false)
	assertInteger(t, 51, 50, 50, false)
	assertInteger(t, 51, 49, 50, false)
	assertInteger(t, 49, 51, 50, true)
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
