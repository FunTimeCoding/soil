package chat_response

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/ollama/ollama/api"
	"testing"
)

func TestResponse(t *testing.T) {
	assert.NotNil(t, New(&api.ChatResponse{}))
}
