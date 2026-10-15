package unit

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/unit/service_tester"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/web"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
	"testing"
)

func newNetboxServer(
	t *testing.T,
	c *mock_client.Client,
	r *memory.Memory,
) *model_context_server.Server {
	t.Helper()
	result := model_context_server.New(
		t,
		func(
			_ *http.ServeMux,
			g *guard.Mux,
		) {
			gonetboxd.Mount(
				c,
				service_tester.Service(t, c),
				web.New(c),
				r,
				mock_recorder.New(),
				g,
			)
		},
	)
	t.Cleanup(result.Close)

	return result
}
