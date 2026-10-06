package index

import "go/types"

func (w *Workspace) storeExternals(
	p *types.Package,
	seen map[string]bool,
) {
	for _, i := range p.Imports() {
		path := i.Path()

		if seen[path] {
			continue
		}

		seen[path] = true

		if _, workspace := w.graph.Nodes[path]; !workspace {
			w.storeExternal(i)
		}

		w.storeExternals(i, seen)
	}
}
