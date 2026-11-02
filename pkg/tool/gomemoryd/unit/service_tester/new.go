package service_tester

import "testing"

func New(t *testing.T) *Tester {
	t.Helper()

	return newWith(t, "", "")
}
