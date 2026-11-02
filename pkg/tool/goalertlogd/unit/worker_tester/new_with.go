package worker_tester

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/event/notifier"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/unit/service_tester"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/worker"
	"github.com/prometheus/client_golang/prometheus"
	"testing"
	"time"
)

func NewWith(
	t *testing.T,
	path string,
	c *mock_client.Client,
	retention time.Duration,
) *Tester {
	t.Helper()
	r := memory.New()
	y := prometheus.NewRegistry()
	v := service_tester.NewServiceWith(t, path, c, retention, notifier.New())
	w := worker.New(
		v.Record,
		constant.PollInterval,
		logger.New(context.Background()),
		y,
		r,
	)
	t.Cleanup(w.Close)

	return &Tester{
		Service:    v,
		Worker:     w,
		MockClient: c,
		Reporter:   r,
		Registry:   y,
		Path:       path,
		t:          t,
	}
}
