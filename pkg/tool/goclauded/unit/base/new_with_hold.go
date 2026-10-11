package base

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/goclauded"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/service_tester"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/web"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
	"testing"
	"time"
)

func NewWithHold(
	t *testing.T,
	hold time.Duration,
) *Server {
	t.Helper()
	s := service_tester.New(t)
	l := logger.New(t.Context())
	result := &Server{
		Tester: s,
		Server: model_context_server.New(
			t,
			func(
				_ *http.ServeMux,
				g *guard.Mux,
			) {
				goclauded.Mount(
					s.Service,
					web.New(s.Service),
					l,
					memory.New(),
					t.TempDir(),
					t.TempDir(),
					hold,
					mock_recorder.New(),
					g,
				)
			},
		),
	}
	t.Cleanup(result.Close)

	return result
}
