package unit

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/daemon"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"testing"
)

func startServer(
	t *testing.T,
	d *daemon.Daemon,
	r *memory.Memory,
) *model_context_server.Server {
	t.Helper()

	return model_context_server.New(
		t,
		func(g *guard.Mux) {
			d.Mount(r, mock_recorder.New(), g)
		},
	)
}
