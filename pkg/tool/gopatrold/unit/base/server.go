package base

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/metric"
	"github.com/funtimecoding/soil/pkg/patrol/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/gopatrold/daemon"
	"testing"
)

type Server struct {
	Reporter  *memory.Memory
	Metric    *metric.Metric
	names     []string
	instances map[string]*mock_client.Client
	daemon    *daemon.Daemon
	t         *testing.T
}
