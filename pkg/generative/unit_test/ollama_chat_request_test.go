package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/generative/ollama/chat_request"
	"testing"
)

func TestOllamaChatRequest(t *testing.T) {
	assert.NotNil(t, chat_request.New())
}
