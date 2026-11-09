package symbol

import (
	"go/types"
	"golang.org/x/tools/go/packages"
)

func namedTypes(all []*packages.Package) []*types.Named {
	var result []*types.Named
	seen := map[*types.Named]bool{}

	for _, p := range all {
		if p.Types == nil {
			continue
		}

		scope := p.Types.Scope()

		for _, name := range scope.Names() {
			o, isType := scope.Lookup(name).(*types.TypeName)

			if !isType || o.IsAlias() {
				continue
			}

			n, isNamed := o.Type().(*types.Named)

			if !isNamed || seen[n] {
				continue
			}

			seen[n] = true
			result = append(result, n)
		}
	}

	return result
}
