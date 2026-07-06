package alert_enricher

import "github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert"

func (e *Enricher) enrich(a *alert.Alert) {
	for _, n := range e.enrichments {
		a.Validate()

		if n.Name == a.Name {
			a.Entity = n.Entity
			a.Category = n.Category
			a.Tag(n.Tags...)
		}
	}
}
