package alert_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert"
	"testing"
)

func Only(
	t *testing.T,
	name string,
	v []*alert.Alert,
) *alert.Alert {
	t.Helper()
	assert.Count(t, 1, v)
	assert.String(t, name, v[0].Name)

	return v[0]
}
