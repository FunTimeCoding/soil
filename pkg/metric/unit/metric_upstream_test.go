package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/identity"
	"github.com/funtimecoding/soil/pkg/metric"
	"github.com/prometheus/client_golang/prometheus"
	"testing"
)

func TestUpstreamSharesOneHistogramAcrossIdentities(t *testing.T) {
	m := metric.New()
	l := prometheus.Labels{"host": "charlie:80", "code": "200", "method": "GET"}
	m.Upstream(identity.New("alfa", "", "")).With(l).Observe(1)
	m.Upstream(identity.New("bravo", "", "")).With(l).Observe(1)
	families, e := m.Registry().Gather()
	errors.PanicOnError(e)
	assert.Count(t, 1, families)
	assert.String(t, "upstream_request_duration_seconds", families[0].GetName())
	var identities []string

	for _, v := range families[0].GetMetric() {
		for _, p := range v.GetLabel() {
			if p.GetName() == "identity" {
				identities = append(identities, p.GetValue())
			}
		}
	}

	assert.Strings(t, []string{"alfa", "bravo"}, identities)
}
