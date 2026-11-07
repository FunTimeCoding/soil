package base

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/jellyfin/mock_client"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/gojellyfind/unit/daemon_tester"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"testing"
)

func New(t *testing.T) *Server {
	t.Helper()
	c := mock_client.New()
	d := daemon_tester.New(t, c)
	r := memory.New()

	return &Server{
		MockClient: c,
		Reporter:   r,
		ContextServer: model_context_server.New(
			t,
			func(g *guard.Mux) {
				d.Mount(r, mock_recorder.New(), g)
			},
		),
	}
}
