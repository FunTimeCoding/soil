package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/system"
	"testing"
)

func TestAbsolutePath(t *testing.T) {
	assert.True(t, len(system.AbsolutePath(constant.CurrentDirectory)) > 0)
}
