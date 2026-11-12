package index

import (
	"github.com/funtimecoding/soil/pkg/source/constant"
	"go/types"
)

func (w *Workspace) storeExternal(p *types.Package) {
	path := p.Path()
	k := externalKey(path, externalVersion(path, w.graph.Requirements))

	for _, kind := range w.kinds {
		if kind.External == nil {
			continue
		}

		target := w.kindStore(constant.IndexExternalKind, kind)

		if kk := w.kindKey(k, kind); !w.holds(target, path, kk) {
			w.keep(target, path, kk, kind.External(p))
		}
	}
}
