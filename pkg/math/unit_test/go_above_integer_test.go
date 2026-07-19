package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/math/go_above"
	"testing"
)

func TestGoAboveInteger(t *testing.T) {
	goAboveIntegerAssertInteger(t, 52, 51, 50, false)
	goAboveIntegerAssertInteger(t, 51, 50, 50, false)
	goAboveIntegerAssertInteger(t, 51, 49, 50, false)
	goAboveIntegerAssertInteger(t, 49, 51, 50, true)
}

func goAboveIntegerAssertInteger(
	t *testing.T,
	past int,
	now int,
	threshold int,
	expect bool,
) {
	t.Helper()
	assert.Boolean(t, expect, go_above.Integer(past, now, threshold))
}
