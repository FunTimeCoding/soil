package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/atlassian/constant"
	"testing"
)

func TestOpsgenieConstant(t *testing.T) {
	assert.String(t, "OPSGENIE_USER_KEY", constant.OpsgenieUserKeyEnvironment)
}
