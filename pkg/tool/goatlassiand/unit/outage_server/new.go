package outage_server

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/unit/service_tester"
	"github.com/funtimecoding/soil/pkg/web/outage_tester"
	"testing"
)

func New(
	t *testing.T,
	r *memory.Memory,
	peer string,
) (*model_context_server.Server, func() error) {
	t.Helper()
	o := service_tester.New(t, service_tester.WithReporter(r))

	switch peer {
	case constant.JiraPeer:
		o.Jira.Failure = outage_tester.RefusedDial()
	case constant.ConfluencePeer:
		o.Confluence.Failure = outage_tester.RefusedDial()
	}

	d := o.Daemon()
	result := model_context_server.Mount(t, d, r)

	return result, outage_tester.Guard(t, d.Workers(), peer).Unavailable
}
