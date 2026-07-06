package label_filter

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert"
	"github.com/funtimecoding/soil/pkg/strings/upper"
	"testing"
)

func TestFilter(t *testing.T) {
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
	f1 := New(false)
	f1.Keep("Apple")
	assertHasOnly(t, f1.Run(fixture), "Alfa")
	f2 := New(true)
	f2.Drop("Apple")
	assertHasOnly(t, f2.Run(fixture), "Bravo")
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
	f3 := New(false)
	f3.KeepValue("Apple", "Red")
	assertHasOnlyValue(t, f3.Run(fixtureValue), "Alfa", "Red")
	f4 := New(true)
	f4.DropValue("Apple", "Red")
	assertHasOnlyValue(t, f4.Run(fixtureValue), "Alfa", "Green")
}

func assertHasOnly(
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
