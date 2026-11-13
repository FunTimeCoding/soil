package gofix

import (
	"github.com/funtimecoding/soil/pkg/tool/gofix/workspace"
	"golang.org/x/tools/go/packages"
)

func RememberedLoad(
	w *workspace.Workspace,
	directory string,
	patterns []string,
) ([]*packages.Package, bool) {
	loaded, _, okay := w.Remembered(loadKey(directory, patterns))

	return loaded, okay
}
