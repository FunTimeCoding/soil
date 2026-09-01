package matcher

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/prometheus/alertmanager/api/v2/models"
	"regexp"
)

func matchesLabels(
	m *models.Matcher,
	l models.LabelSet,
) bool {
	if false {
		console.Format(
			"Matcher: name:%s value:%s regex:%v equal:%v\n",
			*m.Name,
			*m.Value,
			*m.IsRegex,
			*m.IsEqual,
		)
		console.Format("LabelSet: %+v\n", l)
	}

	// Missing labels are treated as empty strings, matching Alertmanager behavior
	value := l[*m.Name]

	if *m.IsRegex {
		r, e := regexp.Compile(*m.Value)

		if e != nil {
			return false
		}

		if *m.IsEqual {
			// =~
			return r.MatchString(value)
		}

		// !~
		return !r.MatchString(value)
	}

	if *m.IsEqual {
		// =
		return value == *m.Value
	}

	// !=
	return value != *m.Value
}
