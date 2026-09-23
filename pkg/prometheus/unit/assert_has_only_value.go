package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert"
	"testing"
)

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
