package fix_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/system"
	"os"
	"path/filepath"
	"testing"
)

func GoFiles(
	t *testing.T,
	directory string,
) map[string]string {
	t.Helper()
	entries, e := os.ReadDir(directory)
	assert.FatalOnError(t, e)
	result := make(map[string]string)

	for _, entry := range entries {
		if !system.IsGoFile(entry.Name()) {
			continue
		}

		result[entry.Name()] = testutil.ReadFile(
			t,
			filepath.Join(directory, entry.Name()),
		)
	}

	return result
}
