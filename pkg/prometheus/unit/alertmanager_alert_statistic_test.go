package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert/statistic"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert/statistic/count"
	"github.com/funtimecoding/soil/pkg/prometheus/constant"
	"testing"
)

func TestCountStatistics(t *testing.T) {
	s := statistic.New()
	assert.Any(
		t,
		&statistic.Statistic{Total: 1},
		s.CountBeforeProcessing(
			[]*alert.Alert{
				{
					State:    constant.ActiveState,
					Severity: constant.CriticalSeverity,
				},
			},
		),
	)
	assert.Any(
		t,
		&statistic.Statistic{
			Total:    1,
			Relevant: 1,
			Severity: count.Severity{Critical: 1},
			State:    count.State{Active: 1},
			Group:    count.Group{All: 1, Other: 1},
		},
		s.CountAfterProcessing(
			[]*alert.Alert{
				{
					State:    constant.ActiveState,
					Severity: constant.CriticalSeverity,
				},
			},
		),
	)
}
