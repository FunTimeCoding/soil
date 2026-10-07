package index

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/source/constant"
	"github.com/funtimecoding/soil/pkg/source/index/record"
	"github.com/funtimecoding/soil/pkg/source/index/xref"
	"github.com/funtimecoding/soil/pkg/source/resolve"
	"slices"
)

func (w *Workspace) refreshReferences() map[string]*record.References {
	g := w.graph
	result := make(map[string]*record.References, len(g.Units))
	keys := make(map[string]string)
	var directories []string

	for path, u := range g.Units {
		k := referencesKey(g, u, w.fileHash(u.Files), w.fingerprints)
		r, found := w.fetch(
			constant.IndexReferencesKind,
			path,
			k,
			record.NewReferences(),
		)

		if found {
			result[path] = r.(*record.References)

			continue
		}

		keys[path] = k
		directories = append(directories, u.Directory)
	}

	if len(keys) == 0 {
		return result
	}

	loaded, e := resolve.LoadTested(
		w.root,
		slices.Compact(slices.Sorted(slices.Values(directories)))...)
	errors.PanicOnError(e)
	workspace := make(map[string]bool, len(g.Nodes))

	for path := range g.Nodes {
		workspace[path] = true
	}

	for _, p := range loaded {
		k, wanted := keys[p.PkgPath]

		if !wanted || len(p.Errors) > 0 || p.TypesInfo == nil {
			continue
		}

		r := xref.Extract(p, workspace)
		w.keep(constant.IndexReferencesKind, p.PkgPath, k, r)
		result[p.PkgPath] = r
	}

	return result
}
