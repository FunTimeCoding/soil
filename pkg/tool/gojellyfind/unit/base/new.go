package base

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/jellyfin/mock_client"
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
		t:        t,
	}
	result.source = result.Source

	for _, f := range o {
		f(result)
	}

	return result
}
