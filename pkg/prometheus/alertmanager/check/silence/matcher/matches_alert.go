package matcher

import (
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/silence"
)

func matchesAlert(
	s *silence.Silence,
	a *alert.Alert,
) bool {
	for _, m := range s.Raw.Matchers {
		if !matchesLabels(m, a.Raw.Labels) {
			return false
		}
	}

	return true
}
