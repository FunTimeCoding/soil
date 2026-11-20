package unit

import (
	"github.com/funtimecoding/soil/pkg/math/unit/math_tester"
	"testing"
)

func TestGoAboveFloat(t *testing.T) {
	m := math_tester.New(t)
	m.AssertGoAboveFloat(false, 52, 51, 50)
	m.AssertGoAboveFloat(false, 51, 50, 50)
	m.AssertGoAboveFloat(false, 51, 49, 50)
	m.AssertGoAboveFloat(true, 49, 51, 50)
}
