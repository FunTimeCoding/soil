package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/generative/ollama/decide_action_prompt"
	"testing"
)

func TestOllamaDecideActionPrompt(t *testing.T) {
	assert.NotNil(t, decide_action_prompt.New())
}
