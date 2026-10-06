package unit

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func treeOf(
	t *testing.T,
	directory string,
) string {
	t.Helper()
	var result strings.Builder
	e := filepath.WalkDir(
		directory,
		func(
			path string,
			d fs.DirEntry,
			f error,
		) error {
			if f != nil || d.IsDir() {
				return f
			}

			content, g := os.ReadFile(path)

			if g != nil {
				return g
			}

			relative, h := filepath.Rel(directory, path)

			if h != nil {
				return h
			}

			_, _ = fmt.Fprintf(&result, "== %s\n%s", relative, content)

			return nil
		},
	)
	assert.FatalOnError(t, e)

	return result.String()
}
