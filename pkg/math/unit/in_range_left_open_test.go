package unit

import (
	"github.com/funtimecoding/soil/pkg/math/ranges"
	"github.com/funtimecoding/soil/pkg/math/unit/math_tester"
	"testing"
)

func TestLeftOpen(t *testing.T) {
	zeroToOne := ranges.Range{L: 0, R: 1}
	math_tester.AssertLeftOpen(t, false, 0, zeroToOne)
	math_tester.AssertLeftOpen(t, true, 0.01, zeroToOne)
	math_tester.AssertLeftOpen(t, true, 0.99, zeroToOne)
	math_tester.AssertLeftOpen(t, true, 1, zeroToOne)
	math_tester.AssertLeftOpen(t, false, 1.01, zeroToOne)
}
