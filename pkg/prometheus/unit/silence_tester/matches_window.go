package silence_tester

import (
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/check/silence/matcher"
	"github.com/prometheus/alertmanager/api/v2/models"
	"time"
)

func MatchesWindow(
	start time.Time,
	end time.Time,
	now time.Time,
) bool {
	s := NewSilence(&start, &end, "alertname", "HighCPU", new(true), new(false))
	a := alert.NewFromLabels(models.LabelSet{"alertname": "HighCPU"})

	return len(matcher.Matches(s, []*alert.Alert{a}, now)) > 0
}
