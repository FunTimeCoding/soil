package base

import (
	"github.com/funtimecoding/soil/pkg/directory/mock_client"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/unit/daemon_tester"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"testing"
)

func NewWith(
	t *testing.T,
	c *mock_client.Client,
	r *memory.Memory,
) *Server {
	t.Helper()
	d := daemon_tester.New(t, c, r)
	result := &Server{
		Reporter: r,
		Server: model_context_server.New(
			t,
			func(g *guard.Mux) {
				d.Mount(r, mock_recorder.New(), g)
			},
			model_context_server.WithSignOn(),
		),
	}
	t.Cleanup(result.Close)

	return result
}
