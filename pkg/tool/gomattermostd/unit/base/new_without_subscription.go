package base

import (
	"net/http"
	"testing"
)

func NewWithoutSubscription(
	t *testing.T,
	configure func(*http.ServeMux),
) *Server {
	t.Helper()

	return serve(t, configure, nil)
}
