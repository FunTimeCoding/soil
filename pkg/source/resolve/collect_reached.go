package resolve

import "go/types"

func collectReached(
	p *types.Package,
	modules []string,
	seen map[string]bool,
	reached map[string]bool,
) {
	if seen[p.Path()] {
		return
	}

	seen[p.Path()] = true

	if inModules(p.Path(), modules) {
		reached[p.Path()] = true
	}

	for _, i := range p.Imports() {
		collectReached(i, modules, seen, reached)
	}
}
