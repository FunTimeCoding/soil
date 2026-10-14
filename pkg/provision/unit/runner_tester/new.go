package runner_tester

import "testing"

func New(t *testing.T) *Tester {
	t.Helper()

	return build(t, func() {}, nil)
}
