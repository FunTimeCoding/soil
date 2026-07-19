package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings/separator"
	"testing"
)

func TestConstant(t *testing.T) {
	assert.String(t, "#", separator.Hash)
	assert.String(t, "//", separator.DoubleSlash)
}
