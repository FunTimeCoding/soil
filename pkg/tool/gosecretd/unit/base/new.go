package base

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/openbao/mock_client"
	"github.com/prometheus/client_golang/prometheus"
	"testing"
)

func New(
	t *testing.T,
	o ...Option,
) *Server {
	t.Helper()
	result := &Server{
		Source:   mock_client.New(t),
		Reporter: memory.New(),
		Registry: prometheus.NewRegistry(),
		t:        t,
	}

	for _, f := range o {
		f(result)
	}

	return result
}
