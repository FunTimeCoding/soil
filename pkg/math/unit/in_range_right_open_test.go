package unit

import (
	"github.com/funtimecoding/soil/pkg/math/ranges"
	"github.com/funtimecoding/soil/pkg/math/unit/math_tester"
	"testing"
)

func TestRightOpen(t *testing.T) {
	zeroToOne := ranges.Range{L: 0, R: 1}
	math_tester.AssertRightOpen(t, true, 0, zeroToOne)
	math_tester.AssertRightOpen(t, true, 0.01, zeroToOne)
	math_tester.AssertRightOpen(t, true, 0.99, zeroToOne)
	math_tester.AssertRightOpen(t, false, 1, zeroToOne)
	math_tester.AssertRightOpen(t, false, 1.01, zeroToOne)
}
