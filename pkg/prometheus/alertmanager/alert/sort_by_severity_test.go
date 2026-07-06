package alert

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/constant"
	"testing"
	"time"
)

func TestSortBySeverity(t *testing.T) {
	assert.Count(
		t,
		1,
		SortBySeverity(
			[]*Alert{
				{Severity: constant.WarningSeverity, Start: new(time.Now())},
			},
			constant.Severities,
		),
	)
}
