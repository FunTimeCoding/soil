package index

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/source/constant"
	"github.com/funtimecoding/soil/pkg/source/resolve"
	"golang.org/x/tools/go/packages"
)

func (w *Workspace) compute(
	paths []string,
	files map[string]string,
	done map[string]bool,
	failed map[string]bool,
) {
	g := w.graph
	loaded, e := resolve.LoadBase(w.root, paths...)
	errors.PanicOnError(e)
	byPath := make(map[string]*packages.Package, len(loaded))

	for _, p := range loaded {
		byPath[p.PkgPath] = p
	}

	seen := make(map[string]bool)

	for _, path := range g.Order {
		if done[path] || failed[path] {
			continue
		}

		p, wanted := byPath[path]

		if !wanted {
			continue
		}

		if len(p.Errors) > 0 || p.Types == nil || !ready(g, path, w.fingerprints) {
			failed[path] = true

			continue
		}

		k := key(g, g.Nodes[path], files[path], w.fingerprints)
		fingerprint := Fingerprint(p.Types)
		w.keep(
			constant.IndexPackageKind,
			path,
			k,
			NewPackageRecord(fingerprint),
		)
		w.fingerprints[path] = fingerprint

		for _, kind := range w.kinds {
			if _, known := w.facts[kind.Name][path]; known {
				continue
			}

			v := kind.Extract(p)
			w.keep(
				kindStore(constant.IndexFactsKind, kind),
				path,
				kindKey(k, kind),
				v,
			)
			w.facts[kind.Name][path] = v
		}

		done[path] = true
		w.storeExternals(p.Types, seen)
	}

	for _, path := range paths {
		if !done[path] {
			failed[path] = true
		}
	}
}
