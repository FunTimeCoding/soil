package assert

import "testing"

func LessEqual(
	t *testing.T,
	than float64,
	actual float64,
) {
	if actual > than {
		t.Helper()
		t.Errorf("\nExpect less equal than: %f\nActual: %f", than, actual)
	}
}
