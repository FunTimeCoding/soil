package name_filter

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert"
	"github.com/funtimecoding/soil/pkg/strings/upper"
	"testing"
)

func TestFilter(t *testing.T) {
	fixture := []*alert.Alert{{Name: upper.Alfa}, {Name: upper.Bravo}}
	f1 := New(false)
	f1.Keep(upper.Alfa)
	assertHasOnly(t, f1.Run(fixture), "Alfa")
	f2 := New(true)
	f2.Drop(upper.Alfa)
	assertHasOnly(t, f2.Run(fixture), "Bravo")
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
