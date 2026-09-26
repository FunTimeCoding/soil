package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/system/constant"
	"strings"
	"testing"
)

func TestWorkDirectory(t *testing.T) {
	assert.True(
		t,
		strings.Contains(system.WorkDirectory(), constant.SystemPackage),
	)
}
