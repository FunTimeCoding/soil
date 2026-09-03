package base

import (
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/relational/lite"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/model_context/mock_recorder"
	generated "github.com/funtimecoding/soil/pkg/tool/gomaintlogd/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/model_context"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/server"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/store"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/web"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
	"testing"
)

func New(t *testing.T) *Server {
	t.Helper()
	s := store.New(lite.NewMemory())
	r := memory.New()

	return &Server{
		Store: s,
		server: model_context_server.New(
			t,
			func(m *http.ServeMux, g *guard.Mux) {
				generated.HandlerFromMux(
					generated.NewStrictHandler(server.New(s, r), nil),
					m,
				)
				model_context.New(
					s,
					r,
					mock_recorder.New(),
					constant.DefaultVersion,
				).Mount(g)
				web.New(s).Mount(g)
			},
		),
	}
}
