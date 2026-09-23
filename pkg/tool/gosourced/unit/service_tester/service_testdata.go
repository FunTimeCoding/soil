package service_tester

import (
	"github.com/funtimecoding/soil/pkg/git"
	"path/filepath"
)

// The fixtures stay beside the service package they exercise;
// the repository root anchors the path from any test directory.
func ServiceTestdata(name string) string {
	return filepath.Join(
		git.FindDirectory(),
		"pkg/tool/gosourced/service/testdata",
		name,
	)
}
