package fix_tester

import "testing"

func New(t *testing.T) *Tester {
	t.Helper()

	return &Tester{t: t}
}
