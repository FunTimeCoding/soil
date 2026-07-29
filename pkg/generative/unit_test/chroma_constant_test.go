package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/generative/constant"
	"testing"
)

func TestChromaConstant(t *testing.T) {
	assert.String(t, "str", constant.ChromaStringType)
	assert.String(t, "int", constant.ChromaIntegerType)
	assert.String(t, "float", constant.ChromaFloatType)
}
