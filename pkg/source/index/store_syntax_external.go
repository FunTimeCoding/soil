package index

import (
	"github.com/funtimecoding/soil/pkg/source/constant"
	"golang.org/x/tools/go/packages"
)

func (w *Workspace) storeSyntaxExternal(p *packages.Package) {
	path := p.Types.Path()
	k := externalKey(path, externalVersion(path, w.graph.Requirements))

	for _, kind := range w.kinds {
		if kind.ExternalSyntax == nil {
			continue
		}

		target := kindStore(constant.IndexExternalKind, kind)

		if kk := w.kindKey(k, kind); !w.holds(target, path, kk) {
			w.keep(target, path, kk, kind.ExternalSyntax(p))
		}
	}
}
