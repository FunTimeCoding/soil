package base

import (
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/relational/lite/connection"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/unit/service_tester"
	"testing"
)

func New(
	t *testing.T,
	option ...service_tester.Option,
) *Server {
	t.Helper()
	o := service_tester.New(t, option...)
	d := connection.New(logger.New(t.Context()), o.Path)
	result := &Server{
		t:        t,
		database: d,
		store:    store.New(d),
		indexer:  o.Indexer,
		Reporter: o.Reporter,
		Server:   model_context_server.Mount(t, o.Daemon(), o.Reporter),
	}
	t.Cleanup(result.Close)

	return result
}
