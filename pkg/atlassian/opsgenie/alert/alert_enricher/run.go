package alert_enricher

import "github.com/funtimecoding/soil/pkg/atlassian/opsgenie/alert"

func (e *Enricher) Run(v []*alert.Alert) []*alert.Alert {
	for _, l := range v {
		e.enrich(l)
	}

	return v
}
