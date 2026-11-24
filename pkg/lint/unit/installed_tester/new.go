package installed_tester

import "testing"

func New(t *testing.T) *Tester {
	t.Helper()

	return &Tester{parent: t.TempDir(), t: t}
}
