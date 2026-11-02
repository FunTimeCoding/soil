package unit

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/goagentd/client"
	"github.com/funtimecoding/soil/pkg/tool/goagentd/daemon"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"net/http"
	"testing"
)

func mountedClient(
	t *testing.T,
	d *daemon.Daemon,
	token string,
) *client.Client {
	t.Helper()
	s := model_context_server.New(
		t,
		func(
			_ *http.ServeMux,
			g *guard.Mux,
		) {
			d.Mount(memory.New(), mock_recorder.New(), g)
		},
	)
	t.Cleanup(s.Close)
	result := client.New(
		locator.New("localhost").Port(s.Port).Insecure(),
		token,
	)
	t.Cleanup(result.Close)

	return result
}
