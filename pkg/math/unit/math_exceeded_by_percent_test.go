package unit

import (
	"github.com/funtimecoding/soil/pkg/math/unit/math_tester"
	"testing"
)

func TestExceededByPercentNoChange(t *testing.T) {
	math_tester.AssertExceededByPercent(t, false, 100, 100, 100, 10)
}

func TestExceededByPercentNinePercent(t *testing.T) {
	math_tester.AssertExceededByPercent(t, false, 100, 109, 100, 10)
}

func TestExceededByPercentTenPercent(t *testing.T) {
	math_tester.AssertExceededByPercent(t, true, 100, 110, 100, 10)
}

func TestExceededByPercentElevenPercent(t *testing.T) {
	math_tester.AssertExceededByPercent(t, true, 100, 111, 100, 10)
}

func TestExceededByPercentNineteenPercent(t *testing.T) {
	math_tester.AssertExceededByPercent(t, true, 100, 119, 100, 10)
}

func TestExceededByPercentTwentyPercent(t *testing.T) {
	math_tester.AssertExceededByPercent(t, true, 100, 120, 100, 10)
}

func TestExceededByPercentTwentyOnePercent(t *testing.T) {
	math_tester.AssertExceededByPercent(t, true, 100, 121, 100, 10)
}

func TestExceededByPercentNinePercentFromHundredTen(t *testing.T) {
	math_tester.AssertExceededByPercent(t, false, 110, 119, 100, 10)
}

func TestExceededByPercentTenPercentFromHundredTen(t *testing.T) {
	math_tester.AssertExceededByPercent(t, true, 110, 120, 100, 10)
}

func TestExceededByPercentElevenPercentFromHundredTen(t *testing.T) {
	math_tester.AssertExceededByPercent(t, true, 110, 121, 100, 10)
}

func TestExceededByPercentRealFigures(t *testing.T) {
	math_tester.AssertExceededByPercent(
		t,
		true,
		661399.2-1,
		661399.2,
		601272,
		10,
	)
}
