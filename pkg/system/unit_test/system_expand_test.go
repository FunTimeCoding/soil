package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/system/constant"
	"testing"
)

func TestExpand(t *testing.T) {
	assert.True(t, len(system.Expand(constant.Tilde)) > 0)
}
