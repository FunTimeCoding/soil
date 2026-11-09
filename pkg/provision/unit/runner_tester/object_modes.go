package runner_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"io/fs"
	"path/filepath"
)

func (o *Tester) ObjectModes() []fs.FileMode {
	o.t.Helper()
	var result []fs.FileMode
	assert.FatalOnError(
		o.t,
		filepath.WalkDir(
			filepath.Join(o.CloneDirectory, ".git", "objects"),
			func(path string, d fs.DirEntry, e error) error {
				if e != nil || d.IsDir() || filepath.Base(filepath.Dir(path)) == "info" {
					return e
				}

				i, f := d.Info()

				if f != nil {
					return f
				}

				result = append(result, i.Mode().Perm())

				return nil
			},
		),
	)
	assert.True(o.t, len(result) > 0)

	return result
}
