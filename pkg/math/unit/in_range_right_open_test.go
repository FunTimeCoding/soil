package unit

import (
	"github.com/funtimecoding/soil/pkg/math/ranges"
	"github.com/funtimecoding/soil/pkg/math/unit/math_tester"
	"testing"
)

func TestRightOpen(t *testing.T) {
	m := math_tester.New(t)
	zeroToOne := ranges.Range{L: 0, R: 1}
	m.AssertRightOpen(true, 0, zeroToOne)
	m.AssertRightOpen(true, 0.01, zeroToOne)
	m.AssertRightOpen(true, 0.99, zeroToOne)
	m.AssertRightOpen(false, 1, zeroToOne)
	m.AssertRightOpen(false, 1.01, zeroToOne)
}
