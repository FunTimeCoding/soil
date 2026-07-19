package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert/alert_enricher/enrichment"
	"github.com/funtimecoding/soil/pkg/strings/upper"
	"testing"
)

func TestEnrichment(t *testing.T) {
	assert.NotNil(t, enrichment.New(upper.Alfa, upper.Bravo, upper.Charlie))
}
