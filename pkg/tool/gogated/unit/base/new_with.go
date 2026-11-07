package base

import (
	"context"
	"github.com/funtimecoding/soil/pkg/directory/face"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/relational"
	"github.com/funtimecoding/soil/pkg/relational/lite"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/gogated/daemon"
	"github.com/funtimecoding/soil/pkg/tool/gogated/store"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"testing"
)

func NewWith(
	t *testing.T,
	d face.DirectoryClient,
) *Server {
	t.Helper()
	o := NewOption(t)
	o.Directory = func() face.DirectoryClient { return d }
	l := logger.New(context.Background())
	r := memory.New()
	v := daemon.New(o, l, r)
	t.Cleanup(v.Close)
	m := lite.New(l, o.LitePath)
	t.Cleanup(func() { relational.CloseMapper(m) })
	s := NewService(t, o, m, d, l, r)
	h := model_context_server.New(
		t,
		func(g *guard.Mux) {
			v.Mount(r, mock_recorder.New(), g)
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
