package outage_server

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/unit/base"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/unit/service_tester"
	"github.com/funtimecoding/soil/pkg/web/outage_tester"
	"testing"
)

func New(
	t *testing.T,
	r *memory.Memory,
) (*model_context_server.Server, func() error) {
	t.Helper()
	s := base.New(t, service_tester.WithReporter(r))
	s.Alertmanager.Fail(outage_tester.RefusedDial())

	return s.Served(), outage_tester.Guard(
		t,
		s.Daemon().Workers(),
		constant.AlertmanagerPeer,
	).Unavailable
}
