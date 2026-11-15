package system_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"os"
	"testing"
)

func ReadOnly(
	t *testing.T,
	directory string,
) {
	t.Helper()
	assert.FatalOnError(t, os.Chmod(directory, 0o500))
	t.Cleanup(
		func() {
			assert.FatalOnError(t, os.Chmod(directory, 0o700))
		},
	)
}
