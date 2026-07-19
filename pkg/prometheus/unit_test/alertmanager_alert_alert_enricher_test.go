package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert/alert_enricher"
	"testing"
)

func TestEnricher(t *testing.T) {
	assert.NotNil(t, alert_enricher.New())
}
