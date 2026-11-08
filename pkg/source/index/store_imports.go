package index

import (
	"github.com/funtimecoding/soil/pkg/source/constant"
	"github.com/funtimecoding/soil/pkg/source/index/record"
	"golang.org/x/tools/go/packages"
	"maps"
	"slices"
)

func (w *Workspace) storeImports(p *packages.Package) {
	k := externalKey(
		p.PkgPath,
		externalVersion(p.PkgPath, w.graph.Requirements),
	)

	if w.holds(constant.IndexImportsKind, p.PkgPath, k) {
		return
	}

	w.keep(
		constant.IndexImportsKind,
		p.PkgPath,
		k,
		record.NewExternal(slices.Sorted(maps.Keys(p.Imports))),
	)
}
