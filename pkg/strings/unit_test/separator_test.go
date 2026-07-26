package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"testing"
)

func TestConstant(t *testing.T) {
	assert.String(t, "#", constant.Hash)
	assert.String(t, "//", constant.DoubleSlash)
}
