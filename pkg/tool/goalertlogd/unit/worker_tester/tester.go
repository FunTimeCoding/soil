package worker_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/service"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/worker"
	"testing"
)

type Tester struct {
	Service    *service.Service
	Worker     *worker.Worker
	MockClient *mock_client.Client
	Reporter   *memory.Memory
	Path       string
	t          *testing.T
}
