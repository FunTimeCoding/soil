package base

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/event/notifier"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/relational/lite"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/gomonitord"
	"github.com/funtimecoding/soil/pkg/tool/gomonitord/store"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
	"testing"
)

func New(t *testing.T) *Server {
	t.Helper()
	events := notifier.New()
	s := store.New(lite.NewMemory(), events)
	result := &Server{
		Store: s,
		Server: model_context_server.New(
			t,
			func(
				_ *http.ServeMux,
				g *guard.Mux,
			) {
				gomonitord.Mount(
					s,
					events,
					memory.New(),
					mock_recorder.New(),
					g,
				)
			},
		),
	}
	t.Cleanup(result.Close)

	return result
}
