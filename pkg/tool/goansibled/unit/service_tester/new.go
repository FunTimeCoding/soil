package service_tester

import (
	"net/http"
	"testing"
)

func New(
	t *testing.T,
	inventory http.Handler,
	option ...Option,
) *Tester {
	t.Helper()
	result := fixture(t, inventory)

	for _, f := range option {
		f(result)
	}

	return result
}
