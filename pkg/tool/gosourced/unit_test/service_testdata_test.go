package unit_test

import (
	"github.com/funtimecoding/soil/pkg/git"
	"path/filepath"
)

// The fixtures stay beside the service package they exercise;
// the repository root anchors the path from any test directory.
func serviceTestdata(name string) string {
	return filepath.Join(
		git.FindDirectory(),
		"pkg/tool/gosourced/service/testdata",
		name,
	)
}
