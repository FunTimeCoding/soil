package math_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/math/normalize_change"
	"testing"
)

func AssertNormalizeChangeInteger(
	t *testing.T,
	expected int,
	now int,
	change int,
	minimum int,
	maximum int,
) {
	t.Helper()
	assert.Integer(
		t,
		expected,
		normalize_change.Integer(now, change, minimum, maximum),
	)
}
