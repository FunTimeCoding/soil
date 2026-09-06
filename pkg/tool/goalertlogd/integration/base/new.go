package base

import (
	"context"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/generative/model_context_server"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/mock_client"
	prometheus "github.com/funtimecoding/soil/pkg/prometheus/constant"
	"github.com/funtimecoding/soil/pkg/relational/lite"
	"github.com/funtimecoding/soil/pkg/telemetry/mock_recorder"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/store"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/web"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/worker"
	"github.com/funtimecoding/soil/pkg/web/guard"
	"net/http"
	"testing"
	"time"
)

func New(t *testing.T) *Server {
	t.Helper()
	s := store.New(lite.NewMemory())
	c := mock_client.New()
	c.Add(
		alert.NewBasic(
			"fp1",
			"HighMemory",
			prometheus.CriticalSeverity,
			"Memory above 90%",
		),
	)
	c.Add(
		alert.NewBasic(
			"fp2",
			"DiskFull",
			prometheus.WarningSeverity,
			"Disk usage above 85%",
		),
	)
	l := logger.New(context.Background())
	r := memory.New()
	w := worker.New(c, s, l, r, 1*time.Minute, 30*24*time.Hour, nil)
	w.Poll()
	v := model_context_server.New(
		t,
		func(_ *http.ServeMux, g *guard.Mux) {
			goalertlogd.Mount(
				s,
				w,
				web.New(s, w),
				r,
				mock_recorder.New(),
				constant.DefaultVersion,
				g,
			)
		},
	)

	return &Server{
		Store:         s,
		Worker:        w,
		MockClient:    c,
		Server: v,
	}
}
