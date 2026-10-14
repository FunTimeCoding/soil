package unclosed_resource

import (
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func (s *Summaries) isReleasing(
	p *packages.Package,
	selector *ast.SelectorExpr,
) bool {
	f, okay := p.TypesInfo.Uses[selector.Sel].(*types.Func)

	return okay && s.holds(methodCondition(f))
}
