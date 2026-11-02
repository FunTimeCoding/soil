package unit

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/client"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/option"
	"github.com/funtimecoding/soil/pkg/tool/goraidd/unit/service_tester"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"net/http"
	"testing"
)

func mountedClient(
	t *testing.T,
	o *option.Raid,
	token string,
) *client.Client {
	t.Helper()
	r := memory.New()
	d := service_tester.NewDaemon(t, o, r)
	s := model_context_server.New(
		t,
		func(_ *http.ServeMux, g *guard.Mux) {
			d.Mount(r, mock_recorder.New(), g)
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
