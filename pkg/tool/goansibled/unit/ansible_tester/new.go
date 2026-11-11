package ansible_tester

import (
	"net/http"
	"testing"
)

func New(
	t *testing.T,
	inventory http.Handler,
) *Tester {
	t.Helper()

	return NewWithSite(t, inventory, "")
}
