package workspace

import (
	"go/token"
	"golang.org/x/tools/go/packages"
)

func (w *Workspace) Remember(
	key string,
	loaded []*packages.Package,
	set *token.FileSet,
) {
	w.loadedKey = key
	w.loadedWrites = w.writes
	w.loaded = loaded
	w.loadedSet = set
}
