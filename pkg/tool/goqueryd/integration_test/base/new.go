package base

import (
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/generative/ollama"
	"github.com/funtimecoding/soil/pkg/relational/lite/connection"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/model_context/mock_recorder"
	generated "github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/model_context"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/server"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/service"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store"
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
		ollama:   l,
		reranker: a,
		server: model_context_server.New(
			t,
			func(m *http.ServeMux) {
				generated.HandlerFromMux(
					generated.NewStrictHandler(
						server.New(v, r),
						nil,
					),
					m,
				)
				model_context.New(
					v,
					r,
					mock_recorder.New(),
					constant.DefaultVersion,
				).Mount(m)
			},
		),
	}
}
