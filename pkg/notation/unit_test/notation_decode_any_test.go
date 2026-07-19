package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/strings/upper"
	"testing"
)

func TestDecodeAny(t *testing.T) {
	var a any
	notation.DecodeAny(true, &a)
	assert.Any(t, a, true)
	notation.DecodeAny(upper.Alfa, &a)
	assert.Any(t, a, "Alfa")
}
