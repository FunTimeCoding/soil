package unit

import (
	"github.com/funtimecoding/soil/pkg/math/unit/math_tester"
	"testing"
)

func TestAboveBelowInteger(t *testing.T) {
	math_tester.AssertAboveBelowInteger(
		t,
		math_tester.AboveBelow{Above: true},
		1,
		0,
	)
	math_tester.AssertAboveBelowInteger(
		t,
		math_tester.AboveBelow{Below: true},
		-1,
		0,
	)
	math_tester.AssertAboveBelowInteger(t, math_tester.AboveBelow{}, 0, 0)
}
