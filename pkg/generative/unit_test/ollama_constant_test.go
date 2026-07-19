package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/generative/ollama/constant"
	"testing"
)

func TestOllamaConstant(t *testing.T) {
	assert.Integer(t, 11434, constant.Port)
	assert.String(t, "llama3.2", constant.Llama32)
	assert.String(t, "llama3.2:1b", constant.Llama321b)
	assert.String(t, "assistant", constant.AssistantRole)
	assert.String(t, "num_ctx", constant.ContextSize)
	assert.String(t, "num_predict", constant.PredictSize)
	assert.String(t, "temperature", constant.Temperature)
}
