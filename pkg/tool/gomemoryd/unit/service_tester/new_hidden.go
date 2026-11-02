package service_tester

import "testing"

func NewHidden(
	t *testing.T,
	tag string,
) *Tester {
	t.Helper()

	return newWith(t, tag, "")
}
