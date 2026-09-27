package assert

import "testing"

func FatalNotNil(
	t *testing.T,
	actual any,
) {
	if nilValue(actual) {
		t.Helper()
		t.Fatalf("expected not nil, got %T", actual)
	}
}
