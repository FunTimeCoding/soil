package unit

import (
	"github.com/funtimecoding/soil/pkg/math/ranges"
	"github.com/funtimecoding/soil/pkg/math/unit/math_tester"
	"testing"
)

func TestOpen(t *testing.T) {
	m := math_tester.New(t)
	zeroToOne := ranges.Range{L: 0, R: 1}
	m.AssertOpen(false, 0, zeroToOne)
	m.AssertOpen(true, 0.01, zeroToOne)
	m.AssertOpen(true, 0.99, zeroToOne)
	m.AssertOpen(false, 1, zeroToOne)
	m.AssertOpen(false, 1.01, zeroToOne)
}
