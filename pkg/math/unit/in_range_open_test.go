package unit

import (
	"github.com/funtimecoding/soil/pkg/math/ranges"
	"github.com/funtimecoding/soil/pkg/math/unit/math_tester"
	"testing"
)

func TestOpen(t *testing.T) {
	zeroToOne := ranges.Range{L: 0, R: 1}
	math_tester.AssertOpen(t, false, 0, zeroToOne)
	math_tester.AssertOpen(t, true, 0.01, zeroToOne)
	math_tester.AssertOpen(t, true, 0.99, zeroToOne)
	math_tester.AssertOpen(t, false, 1, zeroToOne)
	math_tester.AssertOpen(t, false, 1.01, zeroToOne)
}
