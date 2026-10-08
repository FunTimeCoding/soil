package unit

import (
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/check/silence/matcher"
	"github.com/funtimecoding/soil/pkg/prometheus/unit/silence_tester"
	"github.com/prometheus/alertmanager/api/v2/models"
	"time"
)

func matchesSingle(
	name string,
	value string,
	equal bool,
	regex bool,
	labels models.LabelSet,
) bool {
	now := time.Now()
	s := silence_tester.NewSilence(
		new(now.Add(-1*time.Hour)),
		new(now.Add(1*time.Hour)),
		name,
		value,
		new(equal),
		new(regex),
	)

	return len(
		matcher.Matches(s, []*alert.Alert{alert.NewFromLabels(labels)}, now),
	) > 0
}
