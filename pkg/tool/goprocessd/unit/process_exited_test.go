package unit

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/process"
	"golang.org/x/sys/unix"
	"os"
	"testing"
)

func TestGroupAlreadyGoneCountsAsExited(t *testing.T) {
	assert.True(t, process.Exited(unix.ESRCH))
	assert.True(t, process.Exited(fmt.Errorf("getpgid: %w", unix.ESRCH)))
	assert.True(t, process.Exited(os.ErrProcessDone))
}

func TestAnotherSignalFailureIsNotExited(t *testing.T) {
	assert.False(t, process.Exited(unix.EPERM))
}
