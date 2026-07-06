package alert_processor

import "github.com/funtimecoding/soil/pkg/atlassian/opsgenie/alert/alert_enricher"

func New(e *alert_enricher.Enricher) *Processor {
	return &Processor{enricher: e}
}
