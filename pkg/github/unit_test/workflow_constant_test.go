package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/github/constant"
	"testing"
)

func TestWorkflowConstant(t *testing.T) {
	assert.String(t, "active", constant.WorkflowActiveState)
}
