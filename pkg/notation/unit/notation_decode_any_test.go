package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"testing"
)

func TestDecodeAny(t *testing.T) {
	var a any
	notation.DecodeAny(true, &a)
	assert.Any(t, a, true)
	var b any
	notation.DecodeAny(constant.UpperAlfa, &b)
	assert.Any(t, b, "Alfa")
}
