package unclosed_resource

import (
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func registeredCloses(
	p *packages.Package,
	s *ast.SelectorExpr,
	indices map[*types.Var]int,
) []int {
	f, _ := p.TypesInfo.Uses[s.Sel].(*types.Func)

	if !isCloseMethod(f) {
		return nil
	}

	if o := parameter(p, s.X, indices); o != nil {
		return []int{indices[o]}
	}

	return nil
}
