package outage_server

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/goansibled/constant"
	"github.com/funtimecoding/soil/pkg/tool/goansibled/daemon"
	"github.com/funtimecoding/soil/pkg/tool/goansibled/option"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/client"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/face"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"github.com/funtimecoding/soil/pkg/web/outage_tester"
	"github.com/prometheus/client_golang/prometheus"
	"path/filepath"
	"testing"
)

func New(
	t *testing.T,
	r *memory.Memory,
) (*model_context_server.Server, func() error) {
	t.Helper()
	c := client.New(
		locator.New("localhost").Port(1).Insecure(),
		constant.FixtureToken,
	)
	o := option.New()
	o.LitePath = filepath.Join(t.TempDir(), "runs.db")
	o.Inventory = func() face.Client { return c }
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
		join.Colon(constant.InventoryPeer, c.Address()),
	).Unavailable
}
