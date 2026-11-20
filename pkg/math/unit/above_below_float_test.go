package unit

import (
	"github.com/funtimecoding/soil/pkg/math/unit/math_tester"
	"testing"
)

func TestAboveBelowFloat(t *testing.T) {
	m := math_tester.New(t)
	m.AssertAboveBelowFloat(math_tester.AboveBelow{Above: true}, 1, 0)
	m.AssertAboveBelowFloat(math_tester.AboveBelow{Below: true}, -1, 0)
	m.AssertAboveBelowFloat(math_tester.AboveBelow{}, 0, 0)
	m.AssertAboveBelowFloat(math_tester.AboveBelow{}, 1, 1)
	m.AssertAboveBelowFloat(math_tester.AboveBelow{}, -1, 1)
	m.AssertAboveBelowFloat(math_tester.AboveBelow{}, 0, 1)
	m.AssertAboveBelowFloat(math_tester.AboveBelow{Above: true}, 2, 1)
	m.AssertAboveBelowFloat(math_tester.AboveBelow{Below: true}, -2, 1)
}
