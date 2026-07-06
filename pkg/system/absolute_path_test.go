package system

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/constant"
	"testing"
)

func TestAbsolutePath(t *testing.T) {
	assert.True(t, len(AbsolutePath(constant.CurrentDirectory)) > 0)
}
