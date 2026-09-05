package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert"
	prometheus "github.com/funtimecoding/soil/pkg/prometheus/constant"
	strings "github.com/funtimecoding/soil/pkg/strings/constant"
	"testing"
)

func TestFilterSevere(t *testing.T) {
	actual := alert.FilterSevere(
		[]*alert.Alert{
			{
				Name:     strings.UpperAlfa,
				State:    prometheus.ActiveState,
				Severity: prometheus.CriticalSeverity,
			},
			{
				Name:     strings.UpperBravo,
				State:    prometheus.SuppressedState,
				Severity: prometheus.CriticalSeverity,
			},
			{
				Name:     strings.UpperCharlie,
				State:    prometheus.ActiveState,
				Severity: prometheus.InformationSeverity,
			},
		},
	)
	assert.Count(t, 1, actual)
	assert.String(t, "Alfa", actual[0].Name)
}
