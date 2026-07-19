package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert/label_filter"
	"github.com/funtimecoding/soil/pkg/strings/upper"
	"testing"
)

func TestAlertmanagerAlertLabelFilter(t *testing.T) {
	fixture := []*alert.Alert{
		{
			Name:   upper.Alfa,
			Labels: map[string]string{"Apple": "Red"},
		},
		{
			Name:   upper.Bravo,
			Labels: map[string]string{"Banana": "Yellow"},
		},
	}
	f1 := label_filter.New(false)
	f1.Keep("Apple")
	alertmanagerAlertLabelFilterAssertHasOnly(t, f1.Run(fixture), "Alfa")
	f2 := label_filter.New(true)
	f2.Drop("Apple")
	alertmanagerAlertLabelFilterAssertHasOnly(t, f2.Run(fixture), "Bravo")
	fixtureValue := []*alert.Alert{
		{
			Name:   upper.Alfa,
			Labels: map[string]string{"Apple": "Red"},
		},
		{
			Name:   upper.Alfa,
			Labels: map[string]string{"Apple": "Green"},
		},
	}
	f3 := label_filter.New(false)
	f3.KeepValue("Apple", "Red")
	assertHasOnlyValue(t, f3.Run(fixtureValue), "Alfa", "Red")
	f4 := label_filter.New(true)
	f4.DropValue("Apple", "Red")
	assertHasOnlyValue(t, f4.Run(fixtureValue), "Alfa", "Green")
}

func alertmanagerAlertLabelFilterAssertHasOnly(
	t *testing.T,
	v []*alert.Alert,
	name string,
) {
	t.Helper()
	assert.Count(t, 1, v)
	assert.String(t, name, v[0].Name)
}

func assertHasOnlyValue(
	t *testing.T,
	v []*alert.Alert,
	name string,
	value string,
) {
	t.Helper()
	assert.Count(t, 1, v)
	assert.String(t, name, v[0].Name)
	assert.String(t, value, v[0].Labels["Apple"])
}
