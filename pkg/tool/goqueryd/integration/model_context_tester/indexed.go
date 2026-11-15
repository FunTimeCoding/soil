package model_context_tester

import "testing"

func Indexed(t *testing.T) *Tester {
	t.Helper()
	result := New(t)
	result.IndexFixtures()

	return result
}
