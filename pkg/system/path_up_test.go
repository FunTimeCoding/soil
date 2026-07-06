package system

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/system/constant"
	"runtime"
	"testing"
)

func TestPathUp(t *testing.T) {
	switch runtime.GOOS {
	case constant.Linux:
		assert.String(t, "/a", PathUp("/a/b", 1))
	case constant.Windows:
		assert.String(t, "\\a", PathUp("/a/b", 1))
	}
}
