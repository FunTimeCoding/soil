package message

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestMessage(t *testing.T) {
	assert.String(t, "user", RoleUser)
	assert.String(t, "assistant", RoleAssistant)
}
