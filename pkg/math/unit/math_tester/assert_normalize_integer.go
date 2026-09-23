package math_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/math/normalize"
	"testing"
)

func AssertNormalizeInteger(
	t *testing.T,
	expected int,
	i int,
	minimum int,
	maximum int,
) {
	t.Helper()
	normalize.Integer(&i, minimum, maximum)
	assert.Integer(t, expected, i)
}
