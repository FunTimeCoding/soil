package unit

import (
	"github.com/funtimecoding/soil/pkg/math/unit/math_tester"
	"testing"
)

func TestDistribution(t *testing.T) {
	m := math_tester.New(t)
	m.AssertDistribution(100, 100, 10)
	m.AssertDistribution(1000000, 1000000, 10)
}
