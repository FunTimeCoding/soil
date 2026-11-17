package worker_tester

import (
	"context"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/unit/service_tester"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/worker"
	"github.com/prometheus/client_golang/prometheus"
	"testing"
)

func New(
	t *testing.T,
	o ...service_tester.Option,
) *Tester {
	t.Helper()
	s := service_tester.New(t, o...)
	v := s.Service()
	y := prometheus.NewRegistry()
	w := worker.New(
		v.Record,
		constant.PollInterval,
		logger.New(context.Background()),
		y,
		s.Reporter,
	)
	t.Cleanup(w.Close)

	return &Tester{
		Service:      v,
		Worker:       w,
		Alertmanager: s.Alertmanager,
		Reporter:     s.Reporter,
		Registry:     y,
		Path:         s.Path,
		t:            t,
	}
}
