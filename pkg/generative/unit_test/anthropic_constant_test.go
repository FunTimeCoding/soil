package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/constant"
	"testing"
)

func TestAnthropicConstant(t *testing.T) {
	assert.String(t, "ANTHROPIC_TOKEN", constant.TokenEnvironment)
}
