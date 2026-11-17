package base

import (
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert"
	prometheusConstant "github.com/funtimecoding/soil/pkg/prometheus/constant"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/unit/service_tester"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/worker"
	"testing"
)

func New(
	t *testing.T,
	option ...service_tester.Option,
) *Server {
	t.Helper()
	o := service_tester.New(t, option...)
	o.Alertmanager.Add(
		alert.NewBasic(
			"fp1",
			"HighMemory",
			prometheusConstant.CriticalSeverity,
			"Memory above 90%",
		),
	)
	o.Alertmanager.Add(
		alert.NewBasic(
			"fp2",
			"DiskFull",
			prometheusConstant.WarningSeverity,
			"Disk usage above 85%",
		),
	)
	d := o.Daemon()
	s := model_context_server.Mount(t, d, o.Reporter)
	result := &Server{
		Daemon:       d,
		Alertmanager: o.Alertmanager,
		LitePath:     o.Path,
		t:            t,
		Server:       s,
	}

	for _, w := range d.Workers() {
		if k, okay := w.(*worker.Worker); okay {
			result.Clock = k
		}
	}

	result.Clock.Poll()
	t.Cleanup(result.Close)

	return result
}
