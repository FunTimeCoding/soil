package base

import (
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/model_context/mock_recorder"
	generated "github.com/funtimecoding/soil/pkg/tool/gomemoryd/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/model_context"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/server"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/service"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/mock_indexer"
	"net/http"
	"path/filepath"
	"testing"
)

func New(t *testing.T) *Server {
	t.Helper()
	s := store.New(filepath.Join(t.TempDir(), constant.TestDatabase))
	i := mock_indexer.New()
	v := service.New(s, i, i, i)
	r := memory.New()

	return &Server{
		t:       t,
		store:   s,
		indexer: i,
		server: model_context_server.New(
			t,
			func(m *http.ServeMux) {
				generated.HandlerFromMux(
					generated.NewStrictHandler(server.New(v, r), nil),
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
