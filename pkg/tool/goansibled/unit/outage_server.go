package unit

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/goansibled/constant"
	"github.com/funtimecoding/soil/pkg/tool/goansibled/daemon"
	"github.com/funtimecoding/soil/pkg/tool/goansibled/option"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/client"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/face"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"github.com/prometheus/client_golang/prometheus"
	"path/filepath"
	"testing"
)

func outageServer(
	t *testing.T,
	r *memory.Memory,
) (*model_context_server.Server, func() error) {
	t.Helper()
	o := option.New()
	o.LitePath = filepath.Join(t.TempDir(), "runs.db")
	o.Inventory = func() face.Client {
		return client.New(
			locator.New("localhost").Port(1).Insecure(),
			constant.FixtureToken,
		)
	}
	d := daemon.New(o, logger.New(t.Context()), r, prometheus.NewRegistry())
	t.Cleanup(d.Close)
	workers := d.Workers()
	workers[2].Close()
	result := model_context_server.New(
		t,
		func(g *guard.Mux) { d.Mount(r, mock_recorder.New(), g) },
	)
	t.Cleanup(result.Close)

	return result, workers[0].(*peer.Guard).Unavailable
}
