package unit_test

import (
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert/name_filter"
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"testing"
)

func TestAlertmanagerAlertNameFilter(t *testing.T) {
	fixture := []*alert.Alert{
		{Name: constant.UpperAlfa},
		{Name: constant.UpperBravo},
	}
	f1 := name_filter.New(false)
	f1.Keep(constant.UpperAlfa)
	assertHasOnlyAlert(t, f1.Run(fixture), "Alfa")
	f2 := name_filter.New(true)
	f2.Drop(constant.UpperAlfa)
	assertHasOnlyAlert(t, f2.Run(fixture), "Bravo")
}
