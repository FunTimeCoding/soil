package unit

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/unit/daemon_tester"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/unit/service_tester"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"testing"
)

func outageServer(
	t *testing.T,
	r *memory.Memory,
) (*model_context_server.Server, func() error) {
	t.Helper()
	c := service_tester.NewSource(t)
	c.Failure = refusedDial()
	d := daemon_tester.New(t, c, r)
	result := model_context_server.New(
		t,
		func(g *guard.Mux) { d.Mount(r, mock_recorder.New(), g) },
	)
	t.Cleanup(result.Close)

	return result, d.Workers()[0].(*peer.Guard).Unavailable
}
