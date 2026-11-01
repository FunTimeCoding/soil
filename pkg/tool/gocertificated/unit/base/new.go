package base

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/gitlab/mock_client"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/service"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/unit/store_tester"
	"github.com/funtimecoding/soil/pkg/tool/gocertificated/web"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
	"testing"
)

func New(t *testing.T) *Server {
	t.Helper()
	f := mock_client.New()
	t.Cleanup(f.Close)
	r := memory.New()
	v := service.New(store_tester.NewOption(t, f), logger.New(t.Context()), r)
	result := &Server{
		Service:  v,
		Forge:    f,
		Reporter: r,
		Server: model_context_server.New(
			t,
			func(
				_ *http.ServeMux,
				g *guard.Mux,
			) {
				gocertificated.Mount(
					v,
					web.New(v, r),
					r,
					mock_recorder.New(),
					g,
				)
			},
		),
	}
	t.Cleanup(result.Close)

	return result
}
