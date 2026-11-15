package outage_server

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/utilization/mock_client"
	"github.com/funtimecoding/soil/pkg/generative/face"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/daemon"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/peer_tester"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/store_tester"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"github.com/funtimecoding/soil/pkg/web/outage_tester"
	"github.com/prometheus/client_golang/prometheus"
	"path/filepath"
	"testing"
)

func New(
	t *testing.T,
	r *memory.Memory,
	peer string,
) (*model_context_server.Server, func() error) {
	t.Helper()
	d := peer_tester.New(t)
	t.Cleanup(d.Close)
	d.Reporter = r
	u := mock_client.New()
	t.Cleanup(u.Close)

	switch peer {
	case constant.MemoryPeer:
		d.MemoryPeer.Failure = outage_tester.RefusedDial()
	case constant.QueryPeer:
		d.Query.Failure = outage_tester.RefusedDial()
	case constant.UtilizationPeer:
		u.Failure = outage_tester.RefusedDial()
	}

	path := filepath.Join(t.TempDir(), "goclauded.sqlite")
	s := store_tester.NewOnFile(t, path)
	o := d.Option(path, s.Clock())
	o.UtilizationFallback = true
	o.Utilization = func() face.AnthropicClient { return u }
	g := daemon.New(o, logger.New(t.Context()), r, prometheus.NewRegistry())
	t.Cleanup(g.Close)
	result := model_context_server.New(
		t,
		func(m *guard.Mux) {
			g.Mount(r, constant.ChannelHold, mock_recorder.New(), m)
		},
	)
	t.Cleanup(result.Close)

	return result, outage_tester.Guard(t, g.Workers(), peer).Unavailable
}
