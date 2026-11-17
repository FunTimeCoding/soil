package runner_tester

import "testing"

func NewHeld(t *testing.T) *Tester {
	t.Helper()
	gate := make(chan struct{})
	result := New(t, withHold(func() { <-gate }))
	result.gate = gate
	t.Cleanup(result.Release)

	return result
}
