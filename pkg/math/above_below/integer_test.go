package above_below

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestInteger(t *testing.T) {
	assertInteger(t, 1, 0, true, false)
	assertInteger(t, -1, 0, false, true)
	assertInteger(t, 0, 0, false, false)
}

func assertInteger(
	t *testing.T,
	f int,
	magnitude int,
	expectAbove bool,
	expectBelow bool,
) {
	t.Helper()
	var above bool
	var below bool
	Integer(
		f,
		magnitude,
		func() {
			above = true
		},
		func() {
			below = true
		},
	)
	assert.Boolean(t, expectAbove, above)
	assert.Boolean(t, expectBelow, below)
}
