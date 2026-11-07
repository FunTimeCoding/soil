package base

import (
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/daemon"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/double"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/store_tester"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"github.com/prometheus/client_golang/prometheus"
	"path/filepath"
	"testing"
	"time"
)

func serve(
	t *testing.T,
	hold time.Duration,
	d *double.Set,
) *Server {
	t.Helper()
	path := filepath.Join(t.TempDir(), "goclauded.sqlite")
	s := store_tester.NewOnFile(t, path)
	o := d.Option(path, s.Clock())
	g := daemon.New(
		o,
		logger.New(t.Context()),
		d.Reporter,
		prometheus.NewRegistry(),
	)
	t.Cleanup(g.Close)
	result := &Server{
		Set:    d,
		t:      t,
		daemon: g,
		option: o,
		Store:  s,
		Server: model_context_server.New(
			t,
			func(m *guard.Mux) {
				g.Mount(d.Reporter, hold, mock_recorder.New(), m)
			},
		),
	}
	t.Cleanup(result.Close)

	return result
}
