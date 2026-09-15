package assert

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

// Deep comparison over the exported surface only - unexported
// fields are invisible to the comparator, so black-box tests can
// whole-compare entities without reaching into private state.
func Exported(
	t *testing.T,
	expect any,
	actual any,
) {
	t.Helper()

	if d := cmp.Diff(expect, actual, ignoreUnexported()); d != "" {
		t.Errorf("mismatch (-expect +actual):\n%s", d)
	}
}
