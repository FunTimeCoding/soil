package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func readGoFiles(
	t *testing.T,
	directory string,
) map[string]string {
	t.Helper()
	entries, e := os.ReadDir(directory)
	assert.FatalOnError(t, e)
	result := make(map[string]string)

	for _, entry := range entries {
		if !strings.HasSuffix(entry.Name(), constant.GoExtension) {
			continue
		}

		result[entry.Name()] = testutil.ReadFile(
			t,
			filepath.Join(directory, entry.Name()),
		)
	}

	return result
}
