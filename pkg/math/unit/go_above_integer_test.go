package unit

import (
	"github.com/funtimecoding/soil/pkg/math/unit/math_tester"
	"testing"
)

func TestGoAboveInteger(t *testing.T) {
	m := math_tester.New(t)
	m.AssertGoAboveInteger(false, 52, 51, 50)
	m.AssertGoAboveInteger(false, 51, 50, 50)
	m.AssertGoAboveInteger(false, 51, 49, 50)
	m.AssertGoAboveInteger(true, 49, 51, 50)
}
