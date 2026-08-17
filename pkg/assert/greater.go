package assert

import (
	"cmp"
	"testing"
)

func Greater[T cmp.Ordered](
	t *testing.T,
	than T,
	actual T,
) {
	if actual <= than {
		t.Helper()
		t.Errorf("\nExpect greater than: %v\nActual: %v", than, actual)
	}
}
