package unit

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/gitlab/mock_client"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/service"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/unit/service_tester"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/worker"
	"github.com/prometheus/client_golang/prometheus"
	"testing"
	"time"
)

func newWorker(
	t *testing.T,
	r *memory.Memory,
) (*worker.Worker, *service.Service, map[string]*mock_client.Client) {
	t.Helper()
	s, clients := service_tester.New(
		t,
		constant.FixtureInstance,
		constant.FixtureOtherInstance,
	)

	return worker.New(
		s.Refresh,
		s.Entries,
		time.Hour,
		prometheus.NewRegistry(),
		logger.New(t.Context()),
		r,
	), s, clients
}
