package base

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/gosecretd"
	"github.com/funtimecoding/soil/pkg/tool/gosecretd/mock_client"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
	"testing"
)

func New(t *testing.T) *Server {
	t.Helper()
	c := mock_client.New()
	r := memory.New()

	return &Server{
		MockClient: c,
		Reporter:   r,
		ContextServer: model_context_server.New(
			t,
			func(_ *http.ServeMux, g *guard.Mux) {
				gosecretd.Mount(c, r, mock_recorder.New(), g)
			},
		),
	}
}
