package system

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/system/constant"
	"runtime"
	"testing"
)

func TestAbsoluteName(t *testing.T) {
	switch p := runtime.GOOS; p {
	case constant.Linux:
		assert.True(t, AbsoluteName() != "")
	}
}
