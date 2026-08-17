package assert

import (
	"cmp"
	"testing"
)

func GreaterEqual[T cmp.Ordered](
	t *testing.T,
	than T,
	actual T,
) {
	if actual < than {
		t.Helper()
		t.Errorf("\nExpect greater equal than: %v\nActual: %v", than, actual)
	}
}
