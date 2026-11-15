package upstream_server

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func New(
	t *testing.T,
	o ...Option,
) *Server {
	t.Helper()
	result := &Server{t: t, statuses: []int{http.StatusOK}}

	for _, f := range o {
		f(result)
	}

	result.Server = httptest.NewUnstartedServer(http.HandlerFunc(result.serve))
	result.Config.ConnState = result.countConnection

	if result.secure {
		result.StartTLS()
	} else {
		result.Start()
	}

	t.Cleanup(result.Close)

	return result
}
