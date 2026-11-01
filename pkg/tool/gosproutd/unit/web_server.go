package unit

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/service"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/web"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
	"testing"
)

func webServer(
	t *testing.T,
	s *service.Service,
) *model_context_server.Server {
	t.Helper()
	result := model_context_server.New(
		t,
		func(
			_ *http.ServeMux,
			g *guard.Mux,
		) {
			r := memory.New()
			gosproutd.Mount(s, web.New(s, r), r, mock_recorder.New(), g)
		},
	)
	t.Cleanup(result.Close)

	return result
}
