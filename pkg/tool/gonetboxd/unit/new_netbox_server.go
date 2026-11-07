package unit

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/netbox/face"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"testing"
)

func newNetboxServer(
	t *testing.T,
	c face.NetBoxClient,
	r *memory.Memory,
) *model_context_server.Server {
	t.Helper()
	d := newDaemon(t, c)
	result := model_context_server.New(
		t,
		func(g *guard.Mux) {
			d.Mount(r, mock_recorder.New(), g)
		},
	)
	t.Cleanup(result.Close)

	return result
}
