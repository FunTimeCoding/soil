package unit

import (
	"github.com/funtimecoding/soil/pkg/math/unit/math_tester"
	"testing"
)

func TestNormalizeChangeFloat(t *testing.T) {
	m := math_tester.New(t)
	m.AssertNormalizeChangeFloat(1, 0, 1, 0, 100)
	m.AssertNormalizeChangeFloat(-1, 1, -2, 0, 100)
	m.AssertNormalizeChangeFloat(-1, 1, -3, 0, 100)
	m.AssertNormalizeChangeFloat(-100, 100, -100, 0, 100)
	m.AssertNormalizeChangeFloat(-100, 100, -150, 0, 100)
	m.AssertNormalizeChangeFloat(5, 95, 20, 0, 100)
}

func TestFloatNoMaximum(t *testing.T) {
	m := math_tester.New(t)
	m.AssertNormalizeChangeFloat(1, 0, 1, 0, 0)
	m.AssertNormalizeChangeFloat(-1, 1, -2, 0, 0)
	m.AssertNormalizeChangeFloat(-1, 1, -3, 0, 0)
	m.AssertNormalizeChangeFloat(-100, 100, -100, 0, 0)
	m.AssertNormalizeChangeFloat(-100, 100, -150, 0, 0)
}
