package assert

import "testing"

func Panic(
	t *testing.T,
	f func(),
) {
	if _, okay := recovered(f); !okay {
		t.Helper()
		t.Errorf("\nExpect panic\nActual: returned normally")
	}
}
