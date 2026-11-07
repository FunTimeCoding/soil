package base

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/unit/daemon_tester"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/unit/service_tester"
	"github.com/funtimecoding/soil/pkg/web/authorization/authorization_tester"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"testing"
)

func New(t *testing.T) *Server {
	t.Helper()
	m := memory.New()
	d := daemon_tester.New(t, service_tester.NewSource(), m)
	result := &Server{
		Reporter: m,
		Server: model_context_server.NewSignedOn(
			t,
			authorization_tester.Client,
			func(g *guard.Mux) {
				d.Mount(m, mock_recorder.New(), g)
			},
		),
	}
	t.Cleanup(result.Close)

	return result
}
