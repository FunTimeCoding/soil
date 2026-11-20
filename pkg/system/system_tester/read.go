package system_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"os"
	"path/filepath"
)

func (o *Tester) Read(
	base string,
	path string,
) string {
	o.t.Helper()
	b, e := os.ReadFile(filepath.Join(base, path))
	assert.FatalOnError(o.t, e)

	return string(b)
}
