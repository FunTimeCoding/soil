package outage_server

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/client/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/daemon"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/option"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/unit/service_tester"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"github.com/funtimecoding/soil/pkg/web/outage_tester"
	"testing"
	"time"
)

func New(
	t *testing.T,
	r *memory.Memory,
	peer string,
) (*model_context_server.Server, func() error) {
	t.Helper()
	m := service_tester.NewMattermost(t, "kilo")
	k := mock_client.New()
	t.Cleanup(k.Close)

	switch peer {
	case constant.MattermostPeer:
		m.Failure = outage_tester.RefusedDial()
	case constant.RosterPeer:
		k.Fail(outage_tester.RefusedDial())
	}

	d := daemon.New(
		service_tester.NewOption(
			t,
			m,
			k,
			func(o *option.Mattermost) {
				o.Monitor.Enabled = true
				o.Monitor.Interval = time.Hour
			},
		),
		logger.New(t.Context()),
		r,
	)
	t.Cleanup(d.Close)
	result := model_context_server.New(
		t,
		func(g *guard.Mux) { d.Mount(r, mock_recorder.New(), g) },
	)
	t.Cleanup(result.Close)

	return result, outage_tester.Guard(t, d.Workers(), peer).Unavailable
}
