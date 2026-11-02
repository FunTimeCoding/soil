package fixture

import (
	"os"
	"testing"
)

func OpenHandles(t *testing.T) int {
	t.Helper()
	v, e := os.ReadDir("/dev/fd")

	if e != nil {
		t.Fatal(e)
	}

	return len(v)
}
