package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert/name_filter"
	"github.com/funtimecoding/soil/pkg/strings/upper"
	"testing"
)

func TestAlertmanagerAlertNameFilter(t *testing.T) {
	fixture := []*alert.Alert{{Name: upper.Alfa}, {Name: upper.Bravo}}
	f1 := name_filter.New(false)
	f1.Keep(upper.Alfa)
	alertmanagerAlertNameFilterAssertHasOnly(t, f1.Run(fixture), "Alfa")
	f2 := name_filter.New(true)
	f2.Drop(upper.Alfa)
	alertmanagerAlertNameFilterAssertHasOnly(t, f2.Run(fixture), "Bravo")
}

func alertmanagerAlertNameFilterAssertHasOnly(
	t *testing.T,
	v []*alert.Alert,
	name string,
) {
	t.Helper()
	assert.Count(t, 1, v)
	assert.String(t, name, v[0].Name)
}
