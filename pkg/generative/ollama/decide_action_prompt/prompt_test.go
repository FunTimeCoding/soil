package decide_action_prompt

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestPrompt(t *testing.T) {
	assert.NotNil(t, New())
}
