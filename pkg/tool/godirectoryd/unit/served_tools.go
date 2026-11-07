package unit

import (
	"github.com/funtimecoding/soil/pkg/directory/mock_client"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_client"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/godirectoryd/unit/daemon_tester"
	"github.com/funtimecoding/soil/pkg/web/authorization/authorization_tester"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"testing"
)

func servedTools(
	t *testing.T,
	c *mock_client.Client,
	r *memory.Memory,
) *model_context_client.Client {
	t.Helper()
	d := daemon_tester.New(t, c, r)
	s := model_context_server.NewSignedOn(
		t,
		authorization_tester.Client,
		func(g *guard.Mux) {
			d.Mount(r, mock_recorder.New(), g)
		},
	)
	t.Cleanup(s.Close)
	result := s.Connect(t)
	t.Cleanup(result.Close)

	return result
}
