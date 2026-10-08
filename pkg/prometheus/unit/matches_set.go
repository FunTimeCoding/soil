package unit

import (
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/check/silence/matcher"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/silence"
	"github.com/prometheus/alertmanager/api/v2/models"
	"time"
)

func matchesSet(
	matchers []*models.Matcher,
	labels models.LabelSet,
) bool {
	now := time.Now()
	s := silence.NewFromMatchers(
		new(now.Add(-1*time.Hour)),
		new(now.Add(1*time.Hour)),
		matchers,
	)

	return len(
		matcher.Matches(s, []*alert.Alert{alert.NewFromLabels(labels)}, now),
	) > 0
}
