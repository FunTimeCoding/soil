package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert/statistic"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/constant"
	"testing"
)

func TestCountStatistics(t *testing.T) {
	s := statistic.New()
	assert.Any(
		t,
		&statistic.Statistic{
			Total: 1,
		},
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
			Severity: statistic.SeverityCount{Critical: 1},
			State:    statistic.StateCount{Active: 1},
			Group:    statistic.GroupCount{All: 1, Other: 1},
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
