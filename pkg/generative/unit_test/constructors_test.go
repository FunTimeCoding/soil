package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/generative/ollama/chat_request"
	"github.com/funtimecoding/soil/pkg/generative/ollama/chat_response"
	"github.com/funtimecoding/soil/pkg/generative/ollama/classify_prompt"
	"github.com/funtimecoding/soil/pkg/generative/ollama/decide_action_prompt"
	"github.com/funtimecoding/soil/pkg/generative/ollama/generate_request"
	"github.com/funtimecoding/soil/pkg/generative/ollama/generate_response"
	"github.com/funtimecoding/soil/pkg/generative/open_webui/basic"
	"github.com/ollama/ollama/api"
	"testing"
)

func TestConstructors(t *testing.T) {
	assert.NotNil(t, chat_request.New())
	assert.NotNil(t, chat_response.New(&api.ChatResponse{}))
	assert.NotNil(t, classify_prompt.New())
	assert.NotNil(t, decide_action_prompt.New())
	assert.NotNil(t, generate_request.New())
	assert.NotNil(t, generate_response.New(&api.GenerateResponse{}))
	assert.NotNil(t, basic.New("", ""))
}
