package base

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/unit/daemon_tester"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/unit/service_tester"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
	"testing"
)

func New(t *testing.T) *Server {
	t.Helper()
	m := memory.New()
	d := daemon_tester.New(t, service_tester.NewSource(), m)
	result := &Server{
		Reporter: m,
		Server: model_context_server.New(
			t,
			func(
				_ *http.ServeMux,
				g *guard.Mux,
			) {
				d.Mount(m, mock_recorder.New(), g)
			},
		),
	}
	t.Cleanup(result.Close)

	return result
}
