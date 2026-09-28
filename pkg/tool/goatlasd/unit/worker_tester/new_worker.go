package worker_tester

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/face"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/store"
	"github.com/funtimecoding/soil/pkg/tool/goatlasd/worker"
	"github.com/prometheus/client_golang/prometheus"
	"testing"
)

func NewWorker(
	t *testing.T,
	s *store.Store,
	collectors []face.Collector,
	y *prometheus.Registry,
) *worker.Worker {
	t.Helper()

	return worker.New(
		s,
		collectors,
		nil,
		nil,
		constant.FixtureInterval,
		constant.FixtureRetention,
		logger.New(context.Background()),
		y,
		memory.New(),
	)
}
