package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/system/constant"
	"github.com/prometheus/client_golang/prometheus"
	"testing"
)

func reaped(
	t *testing.T,
	y *prometheus.Registry,
	command string,
) int {
	t.Helper()
	families, e := y.Gather()
	assert.FatalOnError(t, e)

	for _, f := range families {
		if f.GetName() != "reaper_reaped_total" {
			continue
		}

		for _, m := range f.GetMetric() {
			for _, l := range m.GetLabel() {
				if l.GetName() == constant.ReapedCommand &&
					l.GetValue() == command {
					return int(m.GetCounter().GetValue())
				}
			}
		}
	}

	return 0
}
