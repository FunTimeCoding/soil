package math_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/math/go_above"
	"testing"
)

func AssertGoAboveFloat(
	t *testing.T,
	expected bool,
	past float64,
	now float64,
	threshold float64,
) {
	t.Helper()
	assert.Boolean(t, expected, go_above.Float(past, now, threshold))
}
