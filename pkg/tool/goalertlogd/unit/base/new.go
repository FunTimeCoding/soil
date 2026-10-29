package base

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/event/notifier"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/mock_client"
	"github.com/funtimecoding/soil/pkg/prometheus/constant"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd"
	toolConstant "github.com/funtimecoding/soil/pkg/tool/goalertlogd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/unit/service_tester"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/web"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/worker"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
	"testing"
)

func New(t *testing.T) *Server {
	t.Helper()
	c := mock_client.New()
	c.Add(
		alert.NewBasic(
			"fp1",
			"HighMemory",
			constant.CriticalSeverity,
			"Memory above 90%",
		),
	)
	c.Add(
		alert.NewBasic(
			"fp2",
			"DiskFull",
			constant.WarningSeverity,
			"Disk usage above 85%",
		),
	)
	v := service_tester.NewService(t, service_tester.DatabasePath(t), c)
	r := memory.New()
	events := notifier.New()
	w := worker.New(
		v,
		events,
		logger.New(context.Background()),
		r,
		toolConstant.PollInterval,
		nil,
	)
	w.Poll()
	s := model_context_server.New(
		t,
		func(
			_ *http.ServeMux,
			g *guard.Mux,
		) {
			goalertlogd.Mount(
				v,
				w,
				web.New(v, w, events),
				r,
				mock_recorder.New(),
				g,
			)
		},
	)
	result := &Server{Service: v, Worker: w, MockClient: c, Server: s}
	t.Cleanup(result.Close)

	return result
}
