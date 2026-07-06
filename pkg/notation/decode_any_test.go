package notation

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings/upper"
	"testing"
)

func TestDecodeAny(t *testing.T) {
	var a any
	DecodeAny(true, &a)
	assert.Any(t, a, true)
	DecodeAny(upper.Alfa, &a)
	assert.Any(t, a, "Alfa")
}
