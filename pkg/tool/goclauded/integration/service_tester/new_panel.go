package service_tester

import "testing"

func NewPanel(
	t *testing.T,
	identifier string,
) *Tester {
	t.Helper()
	result := New(t)
	result.Store.EnsureSession(identifier)
	result.WriteContextLoadFile(identifier)
	result.Service.EnrichSession(identifier)

	return result
}
