package daemon_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/mock_client"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_client"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"testing"
)

func ServedTools(
	t *testing.T,
	c *mock_client.Client,
	r *memory.Memory,
) *model_context_client.Client {
	t.Helper()
	d := New(t, c)
	s := model_context_server.New(
		t,
		func(g *guard.Mux) {
			d.Mount(r, mock_recorder.New(), g)
		},
	)
	t.Cleanup(s.Close)
	result := s.Connect(t)
	t.Cleanup(result.Close)

	return result
}
