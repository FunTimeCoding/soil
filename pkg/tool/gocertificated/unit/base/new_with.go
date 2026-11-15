package base

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/gitlab/mock_client"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/unit/daemon_tester"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/unit/store_tester"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"testing"
	"time"
)

func NewWith(
	t *testing.T,
	clock func() time.Time,
) *Server {
	t.Helper()
	f := mock_client.New()
	r := memory.New()
	o := store_tester.NewOption(t, f)
	o.Clock = clock
	d := daemon_tester.New(t, o, r)
	result := &Server{
		Daemon:   d,
		Forge:    f,
		Reporter: r,
		LitePath: o.LitePath,
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
