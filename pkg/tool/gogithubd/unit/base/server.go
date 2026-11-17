package base

import (
	"github.com/funtimecoding/soil/pkg/metric"
	"github.com/funtimecoding/soil/pkg/tool/gogithubd/daemon"
	"github.com/funtimecoding/soil/pkg/tool/gogithubd/unit/service_tester"
	"testing"
)

type Server struct {
	*service_tester.Tester
	Metric *metric.Metric
	daemon *daemon.Daemon
	t      *testing.T
}
