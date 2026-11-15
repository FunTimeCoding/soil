package unit

import (
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert/name_filter"
	"github.com/funtimecoding/soil/pkg/prometheus/unit/alert_tester"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"testing"
)

func TestAlertmanagerAlertNameFilter(t *testing.T) {
	fixture := []*alert.Alert{
		{Name: stringsConstant.UpperAlfa},
		{Name: stringsConstant.UpperBravo},
	}
	f1 := name_filter.New(false)
	f1.Keep(stringsConstant.UpperAlfa)
	alert_tester.Only(t, "Alfa", f1.Run(fixture))
	f2 := name_filter.New(true)
	f2.Drop(stringsConstant.UpperAlfa)
	alert_tester.Only(t, "Bravo", f2.Run(fixture))
}
