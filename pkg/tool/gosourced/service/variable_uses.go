package service

import (
	"go/ast"
	"go/token"
	"go/types"
	"golang.org/x/tools/go/ast/astutil"
	"golang.org/x/tools/go/packages"
)

func variableUses(
	all []*packages.Package,
	v *types.Var,
	removed []ast.Node,
) (int, int) {
	reads := 0
	writes := 0
	seen := map[token.Pos]bool{}

	for _, p := range all {
		for i, use := range p.TypesInfo.Uses {
			if use.Pos() != v.Pos() ||
				spansContain(removed, i) ||
				seen[i.Pos()] {
				continue
			}

			seen[i.Pos()] = true
			_, file := findOwningFile(all, i.Pos())

			if file == nil {
				continue
			}

			path, _ := astutil.PathEnclosingInterval(file, i.Pos(), i.End())

			if isAssignmentTarget(path) {
				writes++
			} else {
				reads++
			}
		}
	}

	return reads, writes
}
