package silence_tester

import (
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/silence"
	"github.com/prometheus/alertmanager/api/v2/models"
	"time"
)

func NewSilence(
	start *time.Time,
	end *time.Time,
	matcherName string,
	matcherValue string,
	equal *bool,
	regex *bool,
) *silence.Silence {
	m := NewMatcher(matcherName, matcherValue, *equal, *regex)

	return silence.NewFromMatchers(start, end, []*models.Matcher{m})
}
