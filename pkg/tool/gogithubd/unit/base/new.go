package base

import (
	"github.com/funtimecoding/soil/pkg/metric"
	"github.com/funtimecoding/soil/pkg/tool/gogithubd/unit/service_tester"
	"testing"
)

func New(t *testing.T) *Server {
	t.Helper()

	return &Server{
		Tester: service_tester.New(t),
		Metric: metric.New(),
		t:      t,
	}
}
