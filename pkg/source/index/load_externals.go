package index

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"golang.org/x/tools/go/packages"
)

func (w *Workspace) loadExternals(
	paths []string,
	syntax bool,
) {
	mode := packages.NeedName | packages.NeedImports | packages.NeedDeps

	if !syntax {
		mode |= packages.NeedTypes
	}

	closure, e := packages.Load(w.externalConfiguration(mode), paths...)
	errors.PanicOnError(e)
	packages.Visit(
		closure,
		nil,
		func(p *packages.Package) {
			if _, workspace := w.graph.Nodes[p.PkgPath]; workspace {
				return
			}

			w.storeImports(p)

			if !syntax && p.Types != nil {
				w.storeExternal(p.Types)
			}
		},
	)

	if !syntax {
		return
	}

	loaded, e := packages.Load(
		w.externalConfiguration(packages.LoadSyntax),
		paths...,
	)
	errors.PanicOnError(e)

	for _, p := range loaded {
		if p.Types == nil {
			continue
		}

		w.storeExternal(p.Types)

		if p.TypesInfo != nil {
			w.storeSyntaxExternal(p)
		}
	}
}
