package service

import (
	"github.com/funtimecoding/soil/pkg/source/types/resolve_reference"
	"go/token"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func objectReferences(
	all []*packages.Package,
	isTarget func(types.Object) bool,
) []resolve_reference.Reference {
	var result []resolve_reference.Reference
	seen := map[token.Pos]bool{}

	for _, p := range all {
		for ident, use := range p.TypesInfo.Uses {
			if !isTarget(use) || seen[ident.Pos()] {
				continue
			}

			seen[ident.Pos()] = true
			result = append(
				result,
				resolve_reference.Reference{Ident: ident, Package: p},
			)
		}
	}

	return result
}
