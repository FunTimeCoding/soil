package base

import (
	"github.com/funtimecoding/soil/pkg/chat/teams/mock_client"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
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

	for _, f := range o {
		f(result)
	}

	return result
}
