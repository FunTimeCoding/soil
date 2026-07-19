package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/generative/ollama/generate_request"
	"testing"
)

func TestOllamaGenerateRequest(t *testing.T) {
	assert.NotNil(t, generate_request.New())
}
