package convert

import (
	"github.com/funtimecoding/soil/pkg/generative/model_context/paginate"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert/statistic"
)

func AlertResult(
	v []*alert.Alert,
	s *statistic.Statistic,
	limit int,
	offset int,
) *SlimAlertResult {
	return &SlimAlertResult{
		Alerts: paginate.Slice(Alerts(v), limit, offset),
		Total:  s.Total,
		Severity: SlimSeverityCount{
			Critical:    s.Severity.Critical,
			Warning:     s.Severity.Warning,
			Information: s.Severity.Information,
		},
		State: SlimStateCount{
			Active:     s.State.Active,
			Suppressed: s.State.Suppressed,
		},
	}
}
