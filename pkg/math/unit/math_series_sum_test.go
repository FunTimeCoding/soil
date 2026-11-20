package unit

import (
	"github.com/funtimecoding/soil/pkg/math/unit/math_tester"
	"testing"
)

func TestSeriesSum(t *testing.T) {
	m := math_tester.New(t)
	m.AssertSeriesSum(1, 1)
	m.AssertSeriesSum(45, 9)
}
