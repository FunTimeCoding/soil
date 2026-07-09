package base

import (
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/relational/lite"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/model_context/mock_recorder"
	generated "github.com/funtimecoding/soil/pkg/tool/gotelemetryd/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/model_context"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/server"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/service"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/store"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/web"
	"net/http"
	"path/filepath"
	"testing"
)

func New(t *testing.T) *Server {
	t.Helper()
	s := store.New(
		lite.New(filepath.Join(t.TempDir(), constant.TestDatabase)),
	)
	r := memory.New()
	v := model_context_server.New(
		t,
		func(m *http.ServeMux) {
			generated.HandlerFromMux(
				generated.NewStrictHandler(server.New(s, r), nil),
				m,
			)
			model_context.New(
				service.New(s),
				r,
				mock_recorder.New(),
				constant.DefaultVersion,
			).Mount(m)
			web.New(s).Mount(m)
		},
	)

	return &Server{Store: s, ContextServer: v}
}
