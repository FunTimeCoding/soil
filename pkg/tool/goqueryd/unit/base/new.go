package base

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/generative/ollama"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/daemon"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/unit/service_tester"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/unit/shared_reranker"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"github.com/prometheus/client_golang/prometheus"
	"net/http"
	"testing"
)

func New(t *testing.T) *Server {
	t.Helper()
	r := memory.New()
	d := daemon.New(
		service_tester.NewOption(
			t,
			ollama.NewEnvironment(),
			shared_reranker.Lease(),
		),
		logger.New(t.Context()),
		r,
		prometheus.NewRegistry(),
	)
	result := &Server{
		t:      t,
		daemon: d,
		Server: model_context_server.New(
			t,
			func(
				_ *http.ServeMux,
				g *guard.Mux,
			) {
				d.Mount(r, mock_recorder.New(), g)
			},
		),
	}
	t.Cleanup(result.Close)

	return result
}
