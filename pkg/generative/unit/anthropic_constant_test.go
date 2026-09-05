package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/generative/constant"
	"testing"
)

func TestAnthropicConstant(t *testing.T) {
	assert.String(t, "ANTHROPIC_TOKEN", constant.AnthropicTokenEnvironment)
}
