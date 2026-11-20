package fix_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/testutil"
	"github.com/funtimecoding/soil/pkg/system"
	"os"
	"path/filepath"
)

func (o *Tester) GoFiles(directory string) map[string]string {
	o.t.Helper()
	entries, e := os.ReadDir(directory)
	assert.FatalOnError(o.t, e)
	result := make(map[string]string)

	for _, entry := range entries {
		if !system.IsGoFile(entry.Name()) {
			continue
		}

		result[entry.Name()] = testutil.ReadFile(
			o.t,
			filepath.Join(directory, entry.Name()),
		)
	}

	return result
}
