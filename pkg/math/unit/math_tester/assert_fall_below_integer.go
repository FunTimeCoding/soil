package math_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/math/fall_below"
	"testing"
)

func AssertFallBelowInteger(
	t *testing.T,
	expected bool,
	past int,
	now int,
	threshold int,
) {
	t.Helper()
	assert.Boolean(t, expected, fall_below.Integer(past, now, threshold))
}
