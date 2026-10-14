package runner_tester

import (
	"github.com/funtimecoding/soil/pkg/system/run"
	"path/filepath"
)

func (o *Tester) BreakRemote() {
	o.t.Helper()
	c := run.New()
	c.Directory = o.ClonePath
	c.Start(
		"git",
		"config",
		"remote.origin.url",
		filepath.Join(o.t.TempDir(), "missing.git"),
	)
}
