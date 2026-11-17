package worker_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/unit/service_tester"
	"github.com/prometheus/client_golang/prometheus"
	"testing"
)

func New(
	t *testing.T,
	o ...service_tester.Option,
) *Tester {
	t.Helper()

	return &Tester{
		Tester:   service_tester.New(t, o...),
		Registry: prometheus.NewRegistry(),
		t:        t,
	}
}
