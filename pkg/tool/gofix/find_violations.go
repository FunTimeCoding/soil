package gofix

import (
	"go/token"
	"golang.org/x/tools/go/packages"
)

func FindViolations(all []*packages.Package) []violation {
	var result []violation
	seen := make(map[token.Pos]bool)

	for _, p := range all {
		generatedFiles := buildGeneratedSet(p)

		for ident, o := range p.TypesInfo.Defs {
			if o == nil {
				continue
			}

			if seen[o.Pos()] {
				continue
			}

			if generatedFiles[p.Fset.File(ident.Pos()).Name()] {
				continue
			}

			if isInterfaceMethodDefinition(o) {
				continue
			}

			if isInterfaceMethod(p, o) {
				continue
			}

			v := checkNaming(ident, o)

			if v != nil {
				seen[o.Pos()] = true
				result = append(result, *v)
			}
		}
	}

	return result
}
