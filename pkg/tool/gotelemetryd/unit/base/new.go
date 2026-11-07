package base

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/relational"
	"github.com/funtimecoding/soil/pkg/relational/lite"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/daemon"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/store"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"testing"
)

func New(t *testing.T) *Server {
	t.Helper()
	o := NewOption(t)
	l := logger.New(context.Background())
	r := memory.New()
	d := daemon.New(o, l, r)
	t.Cleanup(d.Close)
	m := lite.New(l, o.LitePath)
	t.Cleanup(func() { relational.CloseMapper(m) })
	h := model_context_server.New(
		t,
		func(g *guard.Mux) {
			d.Mount(r, g)
		},
	)
	result := &Server{Store: store.New(m), Daemon: d, Reporter: r, Server: h}
	t.Cleanup(result.Close)

	return result
}
