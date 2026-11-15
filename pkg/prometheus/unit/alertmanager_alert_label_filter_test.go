package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert/label_filter"
	"github.com/funtimecoding/soil/pkg/prometheus/unit/alert_tester"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"testing"
)

func TestAlertmanagerAlertLabelFilter(t *testing.T) {
	fixture := []*alert.Alert{
		{
			Name:   stringsConstant.UpperAlfa,
			Labels: map[string]string{"Apple": "Red"},
		},
		{
			Name:   stringsConstant.UpperBravo,
			Labels: map[string]string{"Banana": "Yellow"},
		},
	}
	f1 := label_filter.New(false)
	f1.Keep("Apple")
	alert_tester.Only(t, "Alfa", f1.Run(fixture))
	f2 := label_filter.New(true)
	f2.Drop("Apple")
	alert_tester.Only(t, "Bravo", f2.Run(fixture))
	fixtureValue := []*alert.Alert{
		{
			Name:   stringsConstant.UpperAlfa,
			Labels: map[string]string{"Apple": "Red"},
		},
		{
			Name:   stringsConstant.UpperAlfa,
			Labels: map[string]string{"Apple": "Green"},
		},
	}
	f3 := label_filter.New(false)
	f3.KeepValue("Apple", "Red")
	assert.String(
		t,
		"Red",
		alert_tester.Only(t, "Alfa", f3.Run(fixtureValue)).Labels["Apple"],
	)
	f4 := label_filter.New(true)
	f4.DropValue("Apple", "Red")
	assert.String(
		t,
		"Green",
		alert_tester.Only(t, "Alfa", f4.Run(fixtureValue)).Labels["Apple"],
	)
}
