package unit_test

import (
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert/name_filter"
	"github.com/funtimecoding/soil/pkg/strings/upper"
	"testing"
)

func TestAlertmanagerAlertNameFilter(t *testing.T) {
	fixture := []*alert.Alert{{Name: upper.Alfa}, {Name: upper.Bravo}}
	f1 := name_filter.New(false)
	f1.Keep(upper.Alfa)
	assertHasOnlyAlert(t, f1.Run(fixture), "Alfa")
	f2 := name_filter.New(true)
	f2.Drop(upper.Alfa)
	assertHasOnlyAlert(t, f2.Run(fixture), "Bravo")
}
