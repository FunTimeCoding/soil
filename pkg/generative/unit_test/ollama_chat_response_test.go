package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/generative/ollama/chat_response"
	"github.com/ollama/ollama/api"
	"testing"
)

func TestOllamaChatResponse(t *testing.T) {
	assert.NotNil(t, chat_response.New(&api.ChatResponse{}))
}
