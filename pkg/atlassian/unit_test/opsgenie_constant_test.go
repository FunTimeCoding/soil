package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/atlassian/opsgenie/constant"
	"testing"
)

func TestOpsgenieConstant(t *testing.T) {
	assert.String(t, "OPSGENIE_USER_KEY", constant.UserKeyEnvironment)
}
