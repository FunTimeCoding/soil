package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/generative/constant"
	"testing"
)

func TestRole(t *testing.T) {
	assert.String(t, "system", constant.RoleSystem)
	assert.String(t, "user", constant.RoleUser)
	assert.String(t, "assistant", constant.RoleAssistant)
}
