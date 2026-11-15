package assert

import (
	systemConstant "github.com/funtimecoding/soil/pkg/system/constant"
	"runtime"
	"testing"
)

func WindowsSkip(t *testing.T) {
	t.Helper()

	if runtime.GOOS == systemConstant.Windows {
		t.Skip("posix fixtures")
	}
}
