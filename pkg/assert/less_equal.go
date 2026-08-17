package assert

import (
	"cmp"
	"testing"
)

func LessEqual[T cmp.Ordered](
	t *testing.T,
	than T,
	actual T,
) {
	if actual > than {
		t.Helper()
		t.Errorf("\nExpect less equal than: %v\nActual: %v", than, actual)
	}
}
