package unit

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/gitlab/mock_client"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/unit/service_tester"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/worker"
	"github.com/prometheus/client_golang/prometheus"
	"testing"
	"time"
)

func newWorker(
	t *testing.T,
	alfa *mock_client.Client,
	bravo *mock_client.Client,
	r *memory.Memory,
) *worker.Worker {
	t.Helper()

	return worker.New(
		service_tester.Service(
			map[string]*mock_client.Client{
				constant.FixtureInstance:      alfa,
				constant.FixtureOtherInstance: bravo,
			},
		),
		time.Hour,
		prometheus.NewRegistry(),
		logger.New(t.Context()),
		r,
	)
}
