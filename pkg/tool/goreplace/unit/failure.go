package unit

import "testing"

func failure(
	t *testing.T,
	e error,
) string {
	t.Helper()

	if e == nil {
		t.Fatal("expected an error, the run succeeded")
	}

	return e.Error()
}
