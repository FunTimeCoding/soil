package assert

import (
	"strings"
	"testing"
)

func StringContains(
	t *testing.T,
	expect string,
	actual string,
) {
	if !strings.Contains(actual, expect) {
		t.Helper()
		t.Errorf("\nExpect string contains: %+q\nActual: %+q", expect, actual)
	}
}
