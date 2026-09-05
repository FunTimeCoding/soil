package integration

import (
	"github.com/funtimecoding/soil/pkg/system/constant"
	"runtime"
	"testing"
)

func windowsSkip(t *testing.T) {
	t.Helper()

	if runtime.GOOS == constant.Windows {
		t.Skip("posix fixtures")
	}
}
