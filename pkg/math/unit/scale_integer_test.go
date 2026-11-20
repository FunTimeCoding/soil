package unit

import (
	"github.com/funtimecoding/soil/pkg/math/unit/math_tester"
	"testing"
)

func TestScaleInteger(t *testing.T) {
	m := math_tester.New(t)
	m.AssertScaleInteger(0, 0, 100, 0)
	m.AssertScaleInteger(50, 0, 100, 0.5)
	m.AssertScaleInteger(100, 0, 100, 1)
	m.AssertScaleInteger(30, 40, 20, 0.5)
	m.AssertScaleInteger(-15, -10, -20, 0.5)
	m.AssertScaleInteger(-20, -30, -10, 0.5)
}
