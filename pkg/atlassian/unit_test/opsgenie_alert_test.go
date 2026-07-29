package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/atlassian/constant"
	"testing"
)

func TestOpsgenieAlert(t *testing.T) {
	assert.Count(t, 2, constant.OpsgenieStatuses)
	assert.Count(t, 5, constant.OpsgeniePriorities)
	assert.Count(t, 1, constant.OpsgenieSkipDetail)
	assert.Count(t, 0, constant.OpsgenieCondenseSkipFields)
}
