package assert

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func Any(
	t *testing.T,
	expect any,
	actual any,
) {
	t.Helper()

	if d := cmp.Diff(expect, actual, exporter()); d != "" {
		t.Errorf("mismatch (-expect +actual):\n%s", d)
	}
}
