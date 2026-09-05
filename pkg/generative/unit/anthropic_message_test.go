package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/generative/constant"
	"testing"
)

func TestMessage(t *testing.T) {
	assert.String(t, "user", constant.AnthropicRoleUser)
	assert.String(t, "assistant", constant.AnthropicRoleAssistant)
}
