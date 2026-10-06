package index

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/source/resolve"
	"golang.org/x/tools/go/packages"
)

func (w *Workspace) loadExternals(paths []string) {
	c := &packages.Config{
		Mode:       packages.NeedName | packages.NeedTypes,
		Dir:        w.root,
		BuildFlags: resolve.BuildFlags(w.root),
	}
	loaded, e := packages.Load(c, paths...)
	errors.PanicOnError(e)
	seen := make(map[string]bool)

	for _, p := range loaded {
		if p.Types == nil {
			continue
		}

		w.storeExternal(p.Types)
		w.storeExternals(p.Types, seen)
	}
}
