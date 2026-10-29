package worker_tester

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/event/notifier"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/service"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/unit/service_tester"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/worker"
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
	o := service_tester.NewOption(path, c)
	o.Retention = retention
	r := memory.New()
	l := logger.New(context.Background())
	v := service.New(o, l, r)
	t.Cleanup(v.Close)
	w := worker.New(v, notifier.New(), l, r, constant.PollInterval, nil)
	t.Cleanup(w.Close)

	return &Tester{
		Service:    v,
		Worker:     w,
		MockClient: c,
		Reporter:   r,
		Path:       path,
		t:          t,
	}
}
