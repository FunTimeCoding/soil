package silence_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager"
	"testing"
)

func Link(
	t *testing.T,
	c *alertmanager.Client,
) string {
	t.Helper()
	v, e := c.Silences(true)
	assert.FatalOnError(t, e)
	assert.Count(t, 1, v)

	return v[0].Link
}
