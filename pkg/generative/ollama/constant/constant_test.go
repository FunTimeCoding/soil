package constant

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestConstant(t *testing.T) {
	assert.Integer(t, 11434, Port)
	assert.String(t, "llama3.2", Llama32)
	assert.String(t, "llama3.2:1b", Llama321b)
	assert.String(t, "assistant", AssistantRole)
	assert.String(t, "num_ctx", ContextSize)
	assert.String(t, "num_predict", PredictSize)
	assert.String(t, "temperature", Temperature)
}
