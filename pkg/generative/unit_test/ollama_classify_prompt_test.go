package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/generative/ollama/classify_prompt"
	"testing"
)

func TestOllamaClassifyPrompt(t *testing.T) {
	assert.NotNil(t, classify_prompt.New())
}
