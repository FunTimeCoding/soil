package base

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/unit/service_tester"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/web"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
	"testing"
)

func New(t *testing.T) *Server {
	t.Helper()
	v := service_tester.Service(t)
	r := memory.New()
	result := &Server{
		Service:  v,
		Reporter: r,
		Server: model_context_server.New(
			t,
			func(
				_ *http.ServeMux,
				g *guard.Mux,
			) {
				gomaintlogd.Mount(v, web.New(v, r), r, mock_recorder.New(), g)
			},
		),
	}
	t.Cleanup(result.Close)

	return result
}
