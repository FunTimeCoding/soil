package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/message"
	"testing"
)

func TestMessage(t *testing.T) {
	assert.String(t, "user", message.RoleUser)
	assert.String(t, "assistant", message.RoleAssistant)
}
