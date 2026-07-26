package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/constant"
	strings "github.com/funtimecoding/soil/pkg/strings/constant"
	"testing"
)

func TestFilterSevere(t *testing.T) {
	actual := alert.FilterSevere(
		[]*alert.Alert{
			{
				Name:     strings.UpperAlfa,
				State:    constant.ActiveState,
				Severity: constant.CriticalSeverity,
			},
			{
				Name:     strings.UpperBravo,
				State:    constant.SuppressedState,
				Severity: constant.CriticalSeverity,
			},
			{
				Name:     strings.UpperCharlie,
				State:    constant.ActiveState,
				Severity: constant.InformationSeverity,
			},
		},
	)
	assert.Count(t, 1, actual)
	assert.String(t, "Alfa", actual[0].Name)
}
