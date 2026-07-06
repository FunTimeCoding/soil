package alert

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/constant"
	"github.com/funtimecoding/soil/pkg/strings/upper"
	"testing"
)

func TestFilterSevere(t *testing.T) {
	actual := FilterSevere(
		[]*Alert{
			{
				Name:     upper.Alfa,
				State:    constant.ActiveState,
				Severity: constant.CriticalSeverity,
			},
			{
				Name:     upper.Bravo,
				State:    constant.SuppressedState,
				Severity: constant.CriticalSeverity,
			},
			{
				Name:     upper.Charlie,
				State:    constant.ActiveState,
				Severity: constant.InformationSeverity,
			},
		},
	)
	assert.Count(t, 1, actual)
	assert.String(t, "Alfa", actual[0].Name)
}
