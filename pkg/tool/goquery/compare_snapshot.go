package goquery

import (
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/terminal"
	"path/filepath"
)

func compareSnapshot(
	t *terminal.Terminal,
	path string,
	body string,
) int {
	stored := snapshotPath(path, false)

	if !system.FileExists(stored) {
		t.Exitf(
			"no snapshot of %s - take one with --snapshot before editing\n",
			path,
		)
	}

	return compareWords(
		system.ReadFile(filepath.Dir(stored), filepath.Base(stored)),
		body,
	)
}
