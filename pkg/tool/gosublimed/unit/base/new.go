package base

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/sublime/mock_client"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/gosublimed/unit/daemon_tester"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"testing"
)

func New(t *testing.T) *Server {
	t.Helper()
	c := mock_client.New()
	d := daemon_tester.New(t, c)
	r := memory.New()
	v := model_context_server.New(
		t,
		func(g *guard.Mux) {
			d.Mount(r, mock_recorder.New(), g)
		},
	)
	result := &Server{MockClient: c, Reporter: r, Server: v}
	t.Cleanup(result.Close)

	return result
}
