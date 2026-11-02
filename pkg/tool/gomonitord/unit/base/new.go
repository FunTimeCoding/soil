package base

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/gomonitord/unit/daemon_tester"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
	"testing"
)

func New(t *testing.T) *Server {
	t.Helper()
	d := daemon_tester.New(t)
	result := &Server{
		Server: model_context_server.New(
			t,
			func(
				_ *http.ServeMux,
				g *guard.Mux,
			) {
				d.Mount(memory.New(), mock_recorder.New(), g)
			},
		),
	}
	t.Cleanup(result.Close)

	return result
}
