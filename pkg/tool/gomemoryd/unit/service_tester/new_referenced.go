package service_tester

import "testing"

func NewReferenced(
	t *testing.T,
	root string,
) *Tester {
	t.Helper()

	return newWith(t, "", root)
}
