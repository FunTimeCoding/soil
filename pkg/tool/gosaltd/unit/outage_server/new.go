package outage_server

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/provision/salt/face"
	"github.com/funtimecoding/soil/pkg/provision/salt/mock_client"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/gosaltd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosaltd/daemon"
	"github.com/funtimecoding/soil/pkg/tool/gosaltd/option"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"github.com/funtimecoding/soil/pkg/web/outage_tester"
	"github.com/prometheus/client_golang/prometheus"
	"path/filepath"
	"testing"
	"time"
)

func New(
	t *testing.T,
	r *memory.Memory,
) (*model_context_server.Server, func() error) {
	t.Helper()
	c := mock_client.New()
	c.Fail(c.Reachable, outage_tester.RefusedDial())
	c.Fail(c.Keys, outage_tester.RefusedDial())
	c.Fail(c.Minions, outage_tester.RefusedDial())
	o := option.New()
	o.LitePath = filepath.Join(t.TempDir(), "gosaltd.db")
	o.KeySyncInterval = 20 * time.Millisecond
	o.Master = func() face.SaltClient { return c }
	d := daemon.New(o, logger.New(t.Context()), r, prometheus.NewRegistry())
	t.Cleanup(d.Close)
	workers := d.Workers()
	workers[len(workers)-1].Close()
	result := model_context_server.New(
		t,
		func(g *guard.Mux) { d.Mount(r, mock_recorder.New(), g) },
	)
	t.Cleanup(result.Close)

	return result, outage_tester.Guard(
		t,
		workers,
		constant.SaltPeer,
	).Unavailable
}
