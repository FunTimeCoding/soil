package gauge_tester

import (
	"github.com/funtimecoding/soil/pkg/metric/metric_tester"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/prometheus/client_golang/prometheus"
	"testing"
)

func Findings(
	t *testing.T,
	r *prometheus.Registry,
	kind string,
) float64 {
	t.Helper()

	return metric_tester.Sample(
		t,
		r,
		"goclauded_findings",
		map[string]string{constant.Kind: kind},
	)
}
