package base

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/tool/gomonitord/unit/daemon_tester"
	"testing"
)

func New(t *testing.T) *Server {
	t.Helper()
	result := &Server{
		Server: model_context_server.Mount(
			t,
			daemon_tester.New(t),
			memory.New(),
		),
	}
	t.Cleanup(result.Close)

	return result
}
