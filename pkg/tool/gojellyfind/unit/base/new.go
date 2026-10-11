package base

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/gojellyfind"
	"github.com/funtimecoding/soil/pkg/tool/gojellyfind/mock_client"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
	"testing"
)

func New(t *testing.T) *Server {
	t.Helper()
	c := mock_client.New()

	return &Server{
		MockClient: c,
		ContextServer: model_context_server.New(
			t,
			func(_ *http.ServeMux, g *guard.Mux) {
				gojellyfind.Mount(c, memory.New(), mock_recorder.New(), g)
			},
		),
	}
}
