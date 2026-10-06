package workspace

import (
	"go/token"
	"golang.org/x/tools/go/packages"
)

func (w *Workspace) Remembered(key string) (
	[]*packages.Package,
	*token.FileSet,
	bool,
) {
	if w.loaded == nil || w.loadedKey != key || w.loadedWrites != w.writes {
		return nil, nil, false
	}

	return w.loaded, w.loadedSet, true
}
