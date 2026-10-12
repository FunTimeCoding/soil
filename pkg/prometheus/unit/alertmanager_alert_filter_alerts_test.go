package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert"
	"github.com/funtimecoding/soil/pkg/prometheus/constant"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"testing"
)

func TestFilterAlerts(t *testing.T) {
	actual := alert.FilterAlerts(
		[]*alert.Alert{
			{
				Name:     stringsConstant.UpperAlfa,
				State:    constant.ActiveState,
				Severity: constant.CriticalSeverity,
			},
			{
				Name:     stringsConstant.UpperBravo,
				State:    constant.SuppressedState,
				Severity: constant.CriticalSeverity,
			},
			{
				Name:     stringsConstant.UpperCharlie,
				State:    constant.ActiveState,
				Severity: constant.InformationSeverity,
			},
		},
	)
	assert.Count(t, 1, actual)
	assert.String(t, "Alfa", actual[0].Name)
}
