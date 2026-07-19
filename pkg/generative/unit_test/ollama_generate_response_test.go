package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/generative/ollama/generate_response"
	"github.com/ollama/ollama/api"
	"testing"
)

func TestOllamaGenerateResponse(t *testing.T) {
	assert.NotNil(t, generate_response.New(&api.GenerateResponse{}))
}
