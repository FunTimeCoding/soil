package base

import (
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/unit/service_tester"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"testing"
)

func New(
	t *testing.T,
	option ...service_tester.Option,
) *Server {
	t.Helper()
	o := service_tester.New(t, option...)
	d := o.Daemon()
	s := model_context_server.New(
		t,
		func(g *guard.Mux) {
			d.Mount(o.Reporter, mock_recorder.New(), g)
		},
	)
	result := &Server{
		Confluence: o.Confluence,
		Jira:       o.Jira,
		Daemon:     d,
		t:          t,
		Server:     s,
	}
	t.Cleanup(result.Close)

	return result
}
