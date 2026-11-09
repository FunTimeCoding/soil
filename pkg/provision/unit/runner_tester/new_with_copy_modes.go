package runner_tester

import "testing"

func NewWithCopyModes(t *testing.T) *Tester {
	t.Helper()

	return build(t, func() {}, nil, nil, true)
}
