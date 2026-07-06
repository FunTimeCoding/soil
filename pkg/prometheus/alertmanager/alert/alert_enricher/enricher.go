package alert_enricher

import "github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert/alert_enricher/enrichment"

type Enricher struct {
	enrichments []*enrichment.Enrichment
}
