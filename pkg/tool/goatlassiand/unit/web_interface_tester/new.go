package web_interface_tester

import "testing"

func New(t *testing.T) *Tester {
	t.Helper()

	return NewWithProject(t, nil)
}
