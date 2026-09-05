package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/generative/constant"
	"testing"
)

func TestN8nConstant(t *testing.T) {
	assert.String(t, "N8N_HOST", constant.N8nHostEnvironment)
}
