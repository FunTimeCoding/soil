package base

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/unit/service_tester"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
	"testing"
)

func New(t *testing.T) *Server {
	t.Helper()
	c := service_tester.NewSource(t)
	v := service_tester.Service(t, c)
	s := model_context_server.New(
		t,
		func(
			_ *http.ServeMux,
			g *guard.Mux,
		) {
			gohabiticad.Mount(v, memory.New(), mock_recorder.New(), g)
		},
	)
	result := &Server{MockClient: c, Server: s}
	t.Cleanup(result.Close)

	return result
}
