package gofix

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/source/resolve"
	"github.com/funtimecoding/soil/pkg/tool/gofix/workspace"
	"go/token"
	"golang.org/x/tools/go/packages"
)

func loadThrough(
	directory string,
	patterns []string,
	w *workspace.Workspace,
) ([]*packages.Package, *token.FileSet) {
	key := loadKey(directory, patterns)

	if loaded, set, okay := w.Remembered(key); okay {
		return loaded, set
	}

	result, set, e := resolve.LoadPackagesThrough(
		directory,
		w.Overlay(),
		patterns...,
	)
	errors.PanicOnError(e)
	w.Remember(key, result, set)

	return result, set
}
