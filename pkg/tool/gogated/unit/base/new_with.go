package base

import (
	"context"
	"github.com/funtimecoding/soil/pkg/directory/face"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/relational"
	"github.com/funtimecoding/soil/pkg/relational/lite"
	"github.com/funtimecoding/soil/pkg/tool/gogated/server"
	"github.com/funtimecoding/soil/pkg/tool/gogated/service"
	"github.com/funtimecoding/soil/pkg/tool/gogated/store"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
	"testing"
)

func NewWith(
	t *testing.T,
	d face.Client,
) *Server {
	t.Helper()
	o := NewOption(t)
	o.Directory = func() face.Client { return d }
	l := logger.New(context.Background())
	r := memory.New()
	s := service.New(o, l, r)
	t.Cleanup(s.Close)
	m := lite.New(l, o.LitePath)
	t.Cleanup(func() { relational.CloseMapper(m) })
	h := model_context_server.New(
		t,
		func(_ *http.ServeMux, g *guard.Mux) {
			server.New(s, r).Mount(g)
		},
	)
	result := &Server{
		Database: m,
		Store:    store.New(m),
		Service:  s,
		Reporter: r,
		Web:      h,
	}
	t.Cleanup(result.Close)

	return result
}
