package unit

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/goalpined/option"
	"github.com/funtimecoding/soil/pkg/tool/goalpined/unit/daemon_tester"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"testing"
)

func mountedReported(
	t *testing.T,
	o *option.Alpine,
	r *memory.Memory,
) *model_context_server.Server {
	t.Helper()
	d := daemon_tester.New(t, o)
	result := model_context_server.New(
		t,
		func(g *guard.Mux) {
			d.Mount(r, mock_recorder.New(), g)
		},
	)
	t.Cleanup(result.Close)

	return result
}
