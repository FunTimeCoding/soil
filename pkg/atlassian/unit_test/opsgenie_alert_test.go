package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/atlassian/opsgenie/alert"
	"testing"
)

func TestOpsgenieAlert(t *testing.T) {
	assert.Count(t, 2, alert.Statuses)
	assert.Count(t, 5, alert.Priorities)
	assert.Count(t, 1, alert.SkipDetail)
	assert.Count(t, 0, alert.CondenseSkipFields)
}
