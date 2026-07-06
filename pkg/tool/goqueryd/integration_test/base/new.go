package base

import (
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/generative/ollama"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/model_context/mock_recorder"
	goqueryd "github.com/funtimecoding/soil/pkg/tool/goqueryd/constant"
	generated "github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/model_context"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/rerank"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/server"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/service"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/store"
	"net/http"
	"path/filepath"
	"testing"
)

func New(t *testing.T) *Server {
	t.Helper()
	s := store.New(filepath.Join(t.TempDir(), constant.TestDatabase))
	l := ollama.NewEnvironment()
	a, e := rerank.New(
		environment.Required(goqueryd.ModelEnvironment),
		environment.Required(goqueryd.TokenizerEnvironment),
	)
	errors.PanicOnError(e)
	t.Cleanup(func() { errors.LogClose(a) })
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
