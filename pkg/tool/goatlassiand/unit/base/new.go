package base

import (
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/unit/service_tester"
	"testing"
)

func New(
	t *testing.T,
	option ...service_tester.Option,
) *Server {
	t.Helper()
	o := service_tester.New(t, option...)
	d := o.Daemon()
	result := &Server{
		Confluence: o.Confluence,
		Jira:       o.Jira,
		Daemon:     d,
		t:          t,
		Server:     model_context_server.Mount(t, d, o.Reporter),
	}
	t.Cleanup(result.Close)

	return result
}
