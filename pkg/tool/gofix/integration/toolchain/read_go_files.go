package toolchain

import (
	"github.com/funtimecoding/soil/pkg/assert"
	soil "github.com/funtimecoding/soil/pkg/constant"
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
		if !strings.HasSuffix(entry.Name(), soil.GoExtension) {
			continue
		}

		result[entry.Name()] = testutil.ReadFile(
			t,
			filepath.Join(directory, entry.Name()),
		)
	}

	return result
}
