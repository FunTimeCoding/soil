package unit

import (
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/check/silence/matcher"
	"github.com/funtimecoding/soil/pkg/prometheus/unit/silence_tester"
	"github.com/prometheus/alertmanager/api/v2/models"
	"time"
)

func matchesWindow(
	start time.Time,
	end time.Time,
	now time.Time,
) bool {
	s := silence_tester.NewSilence(
		&start,
		&end,
		"alertname",
		"HighCPU",
		new(true),
		new(false),
	)
	a := alert.NewFromLabels(models.LabelSet{"alertname": "HighCPU"})

	return len(matcher.Matches(s, []*alert.Alert{a}, now)) > 0
}
