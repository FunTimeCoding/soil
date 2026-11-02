package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert_filter"
	prometheusConstant "github.com/funtimecoding/soil/pkg/prometheus/constant"
	"github.com/funtimecoding/soil/pkg/tool/goalertmanagerd/unit/service_tester"
	"testing"
)

func TestAlertsListSuppressedAndUnratedAlerts(t *testing.T) {
	s := service_tester.NewServing(
		t,
		prometheusConstant.FixtureSuppressedAndUnratedAlerts,
	)
	v, _, e := s.Alerts(t.Context(), "alfa", alert_filter.New())
	assert.FatalOnError(t, e)
	assert.Count(t, 2, v)
}
