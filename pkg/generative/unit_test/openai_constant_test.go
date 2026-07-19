package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/generative/openai/constant"
	"testing"
)

func TestOpenaiConstant(t *testing.T) {
	assert.String(t, "OPENAI_TOKEN", constant.TokenEnvironment)
}
