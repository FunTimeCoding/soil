package alert_enricher

import "github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert/alert_enricher/enrichment"

func (e *Enricher) Add(
	name string,
	entity string,
	category string,
) *Enricher {
	e.enrichments = append(
		e.enrichments,
		enrichment.New(name, entity, category),
	)

	return e
}
