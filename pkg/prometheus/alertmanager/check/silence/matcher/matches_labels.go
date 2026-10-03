package matcher

import (
	"github.com/prometheus/alertmanager/api/v2/models"
	"regexp"
)

func matchesLabels(
	m *models.Matcher,
	l models.LabelSet,
) bool {
	value := l[*m.Name]

	if *m.IsRegex {
		r, e := regexp.Compile(*m.Value)

		if e != nil {
			return false
		}

		if *m.IsEqual {
			return r.MatchString(value)
		}

		return !r.MatchString(value)
	}

	if *m.IsEqual {
		return value == *m.Value
	}

	return value != *m.Value
}
