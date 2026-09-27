package assert

import "testing"

func NotNil(
	t *testing.T,
	actual any,
) {
	if nilValue(actual) {
		t.Helper()
		t.Errorf("expected not nil, got %T", actual)
	}
}
