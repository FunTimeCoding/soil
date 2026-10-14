package unclosed_resource

import (
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func receiverOf(
	p *packages.Package,
	f *ast.FuncDecl,
) *types.Var {
	if f.Recv == nil || len(f.Recv.List) == 0 ||
		len(f.Recv.List[0].Names) == 0 {
		return nil
	}

	result, _ := p.TypesInfo.Defs[f.Recv.List[0].Names[0]].(*types.Var)

	return result
}
