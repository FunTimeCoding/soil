package rewrite_tester

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

func (o *Tester) DirectoryTree(directory string) string {
	o.t.Helper()
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
	assert.FatalOnError(o.t, e)

	return result.String()
}
