package assert

import (
	"github.com/funtimecoding/soil/pkg/math/round"
	"testing"
)

func Round(
	t *testing.T,
	expect float64,
	actual float64,
	decimals int,
) {
	// Pass if identical before rounding
	if actual == expect {
		return
	}

	actual = round.Round(actual, decimals)
	// Pass if identical after rounding
	if actual == expect {
		return
	}

	t.Helper()
	var format string

	switch decimals {
	case 0:
		format = "\nExpect: %.0f\nActual: %.0f"
	case 1:
		format = "\nExpect: %.1f\nActual: %.1f"
	case 2:
		format = "\nExpect: %.2f\nActual: %.2f"
	case 3:
		format = "\nExpect: %.3f\nActual: %.3f"
	case 4:
		format = "\nExpect: %.4f\nActual: %.4f"
	case 5:
		format = "\nExpect: %.5f\nActual: %.5f"
	case 6:
		format = "\nExpect: %.6f\nActual: %.6f"
	case 7:
		format = "\nExpect: %.7f\nActual: %.7f"
	case 8:
		format = "\nExpect: %.8f\nActual: %.8f"
	case 9:
		format = "\nExpect: %.9f\nActual: %.9f"
	case 10:
		format = "\nExpect: %.10f\nActual: %.10f"
	default:
		t.Error("unexpected decimals:", decimals)

		return
	}

	t.Errorf(format, expect, actual)
}
