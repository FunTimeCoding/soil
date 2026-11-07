package base

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/mock_client"
	prometheusConstant "github.com/funtimecoding/soil/pkg/prometheus/constant"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/unit/service_tester"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/worker"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"testing"
)

func New(t *testing.T) *Server {
	t.Helper()
	c := mock_client.New()
	c.Add(
		alert.NewBasic(
			"fp1",
			"HighMemory",
			prometheusConstant.CriticalSeverity,
			"Memory above 90%",
		),
	)
	c.Add(
		alert.NewBasic(
			"fp2",
			"DiskFull",
			prometheusConstant.WarningSeverity,
			"Disk usage above 85%",
		),
	)
	r := memory.New()
	d := service_tester.NewDaemon(
		t,
		service_tester.NewOption(service_tester.DatabasePath(t), c),
		r,
	)
	s := model_context_server.New(
		t,
		func(g *guard.Mux) {
			d.Mount(r, mock_recorder.New(), g)
		},
	)
	result := &Server{Daemon: d, MockClient: c, Server: s}

	for _, w := range d.Workers() {
		if k, okay := w.(*worker.Worker); okay {
			result.Clock = k
		}
	}

	result.Clock.Poll()
	t.Cleanup(result.Close)

	return result
}
