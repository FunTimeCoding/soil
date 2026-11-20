package outage_server

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/client/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/unit/service_tester"
	"github.com/funtimecoding/soil/pkg/web/outage_tester"
	"testing"
)

func New(
	t *testing.T,
	r *memory.Memory,
	peer string,
) (*model_context_server.Server, func() error) {
	t.Helper()
	k := mock_client.New(t)
	o := service_tester.New(
		t,
		service_tester.WithConnector(k),
		service_tester.WithMonitor(),
		service_tester.WithReporter(r),
	)

	switch peer {
	case constant.MattermostPeer:
		o.Mattermost.Failure = outage_tester.RefusedDial()
	case constant.RosterPeer:
		k.Fail(outage_tester.RefusedDial())
	}

	d := o.Daemon()
	result := model_context_server.Mount(t, d, r)

	return result, outage_tester.New(t).Guard(d.Workers(), peer).Unavailable
}
