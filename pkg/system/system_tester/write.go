package system_tester

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"os"
	"path/filepath"
)

func (o *Tester) Write(
	base string,
	path string,
	content string,
) {
	o.t.Helper()
	full := filepath.Join(base, path)
	assert.FatalOnError(o.t, os.MkdirAll(filepath.Dir(full), 0755))
	assert.FatalOnError(o.t, os.WriteFile(full, []byte(content), 0644))
}
