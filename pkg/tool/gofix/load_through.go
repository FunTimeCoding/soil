package gofix

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/source/resolve"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gofix/workspace"
	"go/token"
	"golang.org/x/tools/go/packages"
	"os"
	"slices"
)

func loadThrough(
	directory string,
	patterns []string,
	w *workspace.Workspace,
) ([]*packages.Package, *token.FileSet) {
	key := join.NewLine(slices.Concat([]string{directory}, patterns))

	if loaded, set, okay := w.Remembered(key); okay {
		return loaded, set
	}

	result, set, e := resolve.LoadPackagesThrough(
		directory,
		w.Overlay(),
		patterns...,
	)

	if e != nil {
		errors.Printf("load: %s\n", e)
		os.Exit(1)
	}

	w.Remember(key, result, set)

	return result, set
}
