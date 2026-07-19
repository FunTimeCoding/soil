package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/generative/chroma/constant"
	"testing"
)

func TestChromaConstant(t *testing.T) {
	assert.String(t, "str", constant.Str)
	assert.String(t, "int", constant.Int)
	assert.String(t, "float", constant.Float)
}
