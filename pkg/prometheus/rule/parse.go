package rule

import (
	"github.com/funtimecoding/soil/pkg/prometheus/constant"
	stringLibrary "github.com/funtimecoding/soil/pkg/strings"
	"strings"
)

func (r *Rule) parse() *Rule {
	if r.RawAlert == nil {
		return r
	}

	for k, v := range r.RawAlert.Labels {
		switch k {
		case constant.SeverityKey:
			r.Severity = string(v)
		}
	}

	for k, v := range r.RawAlert.Annotations {
		switch k {
		case constant.SummaryKey:
			r.Summary = string(v)
		case constant.DescriptionKey:
			r.Description = strings.TrimSpace(string(v))
		case constant.DurationKey:
			r.Duration = stringLibrary.MustToInteger(string(v))
		case constant.RunbookKey:
			r.Runbook = string(v)
		case constant.DocumentationKey:
			r.Runbook = string(v)
		}

		if r.Runbook != "" {
			break
		}
	}

	return r
}
