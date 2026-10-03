package assert

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

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
