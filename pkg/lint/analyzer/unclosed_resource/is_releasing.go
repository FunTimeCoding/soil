package unclosed_resource

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func (s *Summaries) isReleasing(
	p *packages.Package,
	selector *ast.SelectorExpr,
) bool {
	f, okay := p.TypesInfo.Uses[selector.Sel].(*types.Func)

	return okay && (constant.ReleasingMethods[functionKey(f)] ||
		s.holds(methodCondition(f)))
}
