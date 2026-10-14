package runner_tester

import "testing"

func NewHeld(t *testing.T) *Tester {
	t.Helper()
	gate := make(chan struct{})
	result := build(t, func() { <-gate }, nil)
	result.gate = gate
	t.Cleanup(result.Release)

	return result
}
