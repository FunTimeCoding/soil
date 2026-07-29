package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/generative/constant"
	"testing"
)

func TestOllamaConstant(t *testing.T) {
	assert.Integer(t, 11434, constant.OllamaPort)
	assert.String(t, "llama3.2", constant.Llama32)
	assert.String(t, "llama3.2:1b", constant.Llama321b)
	assert.String(t, "assistant", constant.OllamaAssistantRole)
	assert.String(t, "num_ctx", constant.OllamaContextSize)
	assert.String(t, "num_predict", constant.OllamaPredictSize)
	assert.String(t, "temperature", constant.OllamaTemperature)
}
