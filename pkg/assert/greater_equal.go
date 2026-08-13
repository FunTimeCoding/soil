package assert

import "testing"

func GreaterEqual(
	t *testing.T,
	than float64,
	actual float64,
) {
	if actual < than {
		t.Helper()
		t.Errorf("\nExpect greater equal than: %f\nActual: %f", than, actual)
	}
}
