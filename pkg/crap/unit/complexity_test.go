package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestComplexityStraightLine(t *testing.T) {
	assert.Integer(t, 1, measure(t, `func f() { a := 1; _ = a }`))
}

func TestComplexityBranches(t *testing.T) {
	assert.Integer(
		t,
		4,
		measure(
			t,
			`func f(a, b bool) {
	if a { }
	for i := 0; i < 2; i++ { }
	for range []int{} { }
}`,
		),
	)
}

func TestComplexitySwitchAndSelect(t *testing.T) {
	assert.Integer(
		t,
		4,
		measure(
			t,
			`func f(a int, c chan int) {
	switch a {
	case 1:
	case 2:
	default:
	}
	select {
	case <-c:
	default:
	}
}`,
		),
	)
}

func TestComplexityBooleanOperators(t *testing.T) {
	assert.Integer(
		t,
		3,
		measure(t, `func f(a, b, c bool) bool { return a && b || c }`),
	)
}
