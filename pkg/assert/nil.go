package assert

import "testing"

func Nil(
	t *testing.T,
	actual any,
) {
	if !nilValue(actual) {
		t.Helper()
		t.Errorf("expected nil, got %T", actual)
	}
}
