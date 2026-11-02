package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	soil "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/system"
	"testing"
)

func TestAbsolutePath(t *testing.T) {
	assert.True(t, len(system.AbsolutePath(soil.CurrentDirectory)) > 0)
}
