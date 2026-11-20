package outage_server

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goalertmanagerd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goalertmanagerd/unit/base"
	"github.com/funtimecoding/soil/pkg/tool/goalertmanagerd/unit/service_tester"
	"github.com/funtimecoding/soil/pkg/web/outage_tester"
	"testing"
)

func New(
	t *testing.T,
	r *memory.Memory,
	peer string,
) (*model_context_server.Server, func() error) {
	t.Helper()
	s := base.New(t, service_tester.WithReporter(r))

	switch peer {
	case constant.AlertmanagerPeer:
		s.Alfa().Alertmanager.Fail(outage_tester.RefusedDial())
	case constant.PrometheusPeer:
		s.Alfa().Prometheus.Failure = outage_tester.RefusedDial()
	}

	return s.Served(), outage_tester.New(t).Guard(
		s.Daemon().Workers(),
		join.Colon(peer, "alfa"),
	).Unavailable
}
