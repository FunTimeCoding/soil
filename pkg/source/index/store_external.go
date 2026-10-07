package index

import (
	"github.com/funtimecoding/soil/pkg/source/constant"
	"github.com/funtimecoding/soil/pkg/source/index/record"
	"go/types"
)

func (w *Workspace) storeExternal(p *types.Package) {
	path := p.Path()
	k := externalKey(path, externalVersion(path, w.graph.Requirements))

	if !w.holds(constant.IndexImportsKind, path, k) {
		var imports []string

		for _, d := range p.Imports() {
			imports = append(imports, d.Path())
		}

		w.keep(constant.IndexImportsKind, path, k, record.NewExternal(imports))
	}

	for _, kind := range w.kinds {
		if kind.External == nil {
			continue
		}

		target := kindStore(constant.IndexExternalKind, kind)

		if kk := kindKey(k, kind); !w.holds(target, path, kk) {
			w.keep(target, path, kk, kind.External(p))
		}
	}
}
