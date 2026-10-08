package gofix

import (
	"github.com/funtimecoding/soil/pkg/lint/face"
	"go/token"
	"golang.org/x/tools/go/packages"
)

func FindViolations(
	all []*packages.Package,
	reported map[string]bool,
	faces *face.Set,
) []Violation {
	var result []Violation
	seen := make(map[token.Pos]bool)

	for _, p := range all {
		if !reported[p.PkgPath] {
			continue
		}

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

			if faces.Implements(o) {
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
