package system

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/system/constant"
	"testing"
)

func TestExpand(t *testing.T) {
	assert.True(t, len(Expand(constant.Tilde)) > 0)
}
