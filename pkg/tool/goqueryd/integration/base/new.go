package base

import (
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/generative/ollama"
	"github.com/funtimecoding/soil/pkg/relational/lite/connection"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/service"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/web"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
	"testing"
)

func New(t *testing.T) *Server {
	t.Helper()
	s := store.New(connection.NewMemory())
	l := ollama.NewEnvironment()
	a := sharedReranker()
	v := service.New(s, l, a)
	r := memory.New()

	return &Server{
		t:        t,
		store:    s,
		embedder: l,
		reranker: a,
		Server: model_context_server.New(
			t,
			func(_ *http.ServeMux, g *guard.Mux) {
				goqueryd.Mount(
					v,
					web.New(v),
					r,
					mock_recorder.New(),
					constant.DefaultVersion,
					g,
				)
			},
		),
	}
}
