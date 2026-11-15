package system_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"os"
	"testing"
)

func Unreadable(
	t *testing.T,
	directory string,
) {
	t.Helper()
	aside := join.Empty(directory, ".aside")
	assert.FatalOnError(t, os.Rename(directory, aside))
	assert.FatalOnError(t, os.WriteFile(directory, nil, 0o644))
	t.Cleanup(
		func() {
			assert.FatalOnError(t, os.Remove(directory))
			assert.FatalOnError(t, os.Rename(aside, directory))
		},
	)
}
