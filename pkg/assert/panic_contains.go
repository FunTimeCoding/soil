package assert

import (
	"fmt"
	"testing"
)

func PanicContains(
	t *testing.T,
	expect string,
	f func(),
) {
	t.Helper()
	v, okay := recovered(f)

	if !okay {
		t.Errorf(
			"\nExpect panic containing: %+q\nActual: returned normally",
			expect,
		)

		return
	}

	StringContains(t, expect, fmt.Sprint(v))
}
