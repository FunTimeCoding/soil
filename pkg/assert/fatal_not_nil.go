package assert

import "testing"

func FatalNotNil(
	t *testing.T,
	actual any,
) {
	if actual == nil {
		t.Helper()
		t.Fatalf("expected not nil, got nil")
	}
}
