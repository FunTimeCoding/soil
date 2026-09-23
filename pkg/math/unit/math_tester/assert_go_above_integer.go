package math_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/math/go_above"
	"testing"
)

func AssertGoAboveInteger(
	t *testing.T,
	expected bool,
	past int,
	now int,
	threshold int,
) {
	t.Helper()
	assert.Boolean(t, expected, go_above.Integer(past, now, threshold))
}
