package unit

import (
	"github.com/funtimecoding/soil/pkg/math/unit/math_tester"
	"testing"
)

func TestScaleFloat(t *testing.T) {
	m := math_tester.New(t)
	m.AssertScaleFloat(0, 0, 100, 0)
	m.AssertScaleFloat(50, 0, 100, 0.5)
	m.AssertScaleFloat(100, 0, 100, 1)
	m.AssertScaleFloat(30, 40, 20, 0.5)
	m.AssertScaleFloat(1, 1, 1, 0.5)
	m.AssertScaleFloat(0, -1, 1, 0.5)
	m.AssertScaleFloat(0.7, 2.0/3.0, 1, 0)
	m.AssertScaleFloat(1, 2.0/3.0, 1, 1)
}
