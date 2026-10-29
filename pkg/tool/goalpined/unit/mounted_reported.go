package unit

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/goalpined"
	"github.com/funtimecoding/soil/pkg/tool/goalpined/service"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
	"testing"
)

func mountedReported(
	t *testing.T,
	v *service.Service,
	r *memory.Memory,
) *model_context_server.Server {
	t.Helper()
	result := model_context_server.New(
		t,
		func(
			_ *http.ServeMux,
			g *guard.Mux,
		) {
			goalpined.Mount(v, r, mock_recorder.New(), g)
		},
	)
	t.Cleanup(result.Close)

	return result
}
