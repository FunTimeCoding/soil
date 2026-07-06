package rule

import (
	stringLibrary "github.com/funtimecoding/soil/pkg/strings"
	"strings"
)

func (r *Rule) parse() *Rule {
	if r.RawAlert == nil {
		return r
	}

	for k, v := range r.RawAlert.Labels {
		switch k {
		case SeverityKey:
			r.Severity = string(v)
		}
	}

	for k, v := range r.RawAlert.Annotations {
		switch k {
		case SummaryKey:
			r.Summary = string(v)
		case DescriptionKey:
			r.Description = strings.TrimSpace(string(v))
		case DurationKey:
			r.Duration = stringLibrary.MustToInteger(string(v))
		case RunbookKey:
			r.Runbook = string(v)
		case DocumentationKey:
			r.Runbook = string(v)
		}

		if r.Runbook != "" {
			break
		}
	}

	return r
}
