package unit

import (
	"github.com/funtimecoding/soil/pkg/math/ranges"
	"github.com/funtimecoding/soil/pkg/math/unit/math_tester"
	"testing"
)

func TestLeftOpen(t *testing.T) {
	m := math_tester.New(t)
	zeroToOne := ranges.Range{L: 0, R: 1}
	m.AssertLeftOpen(false, 0, zeroToOne)
	m.AssertLeftOpen(true, 0.01, zeroToOne)
	m.AssertLeftOpen(true, 0.99, zeroToOne)
	m.AssertLeftOpen(true, 1, zeroToOne)
	m.AssertLeftOpen(false, 1.01, zeroToOne)
}
