package matcher

import (
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/silence"
	"time"
)

func Matches(
	s *silence.Silence,
	v []*alert.Alert,
	now time.Time,
) []*alert.Alert {
	var result []*alert.Alert

	// Silence is active in the range [start, end) - inclusive start, exclusive end
	if now.Before(*s.Start) || !now.Before(*s.End) {
		return result
	}

	for _, a := range v {
		if matchesAlert(s, a) {
			result = append(result, a)
		}
	}

	return result
}
