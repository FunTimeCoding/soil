package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/math/go_above"
	"testing"
)

func TestGoAboveFloat(t *testing.T) {
	goAboveFloatAssertFloat(t, 52, 51, 50, false)
	goAboveFloatAssertFloat(t, 51, 50, 50, false)
	goAboveFloatAssertFloat(t, 51, 49, 50, false)
	goAboveFloatAssertFloat(t, 49, 51, 50, true)
}

func goAboveFloatAssertFloat(
	t *testing.T,
	past float64,
	now float64,
	threshold float64,
	expect bool,
) {
	t.Helper()
	assert.Boolean(t, expect, go_above.Float(past, now, threshold))
}
