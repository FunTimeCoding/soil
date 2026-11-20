package unit

import (
	"github.com/funtimecoding/soil/pkg/math/unit/math_tester"
	"testing"
)

func TestWeight(t *testing.T) {
	m := math_tester.New(t)
	m.AssertWeight(1, 1, 1, 1, 1)
	m.AssertWeight(0.75, 1, 0.5, 1, 1)
	m.AssertWeight(0.5, 1, 0, 1, 1)
	m.AssertWeight(127.5, 255, 0, 1, 1)
	m.AssertWeight(0.67, 1, 0, 2, 1)
	m.AssertWeight(0.33, 1, 0, 1, 2)
}
