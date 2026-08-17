package assert

import (
	"cmp"
	"testing"
)

func Less[T cmp.Ordered](
	t *testing.T,
	than T,
	actual T,
) {
	if actual >= than {
		t.Helper()
		t.Errorf("\nExpect less than: %v\nActual: %v", than, actual)
	}
}
