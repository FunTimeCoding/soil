package base

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/option"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/unit/service_tester"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"testing"
)

func Mount(
	t *testing.T,
	o *option.Raid,
	r *memory.Memory,
) *model_context_server.Server {
	t.Helper()
	d := service_tester.NewDaemon(t, o, r)
	result := model_context_server.New(
		t,
		func(g *guard.Mux) {
			d.Mount(r, mock_recorder.New(), g)
		},
		model_context_server.WithSignOn(),
	)
	t.Cleanup(result.Close)

	return result
}
