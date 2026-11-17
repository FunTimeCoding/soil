package base

import (
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/relational/lite/connection"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/unit/service_tester"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"testing"
)

func New(
	t *testing.T,
	option ...service_tester.Option,
) *Server {
	t.Helper()
	o := service_tester.New(t, option...)
	v := o.Daemon()
	d := connection.New(logger.New(t.Context()), o.Path)
	result := &Server{
		t:        t,
		database: d,
		store:    store.New(d),
		indexer:  o.Indexer,
		Reporter: o.Reporter,
		Server: model_context_server.New(
			t,
			func(g *guard.Mux) {
				v.Mount(o.Reporter, mock_recorder.New(), g)
			},
		),
	}
	t.Cleanup(result.Close)

	return result
}
