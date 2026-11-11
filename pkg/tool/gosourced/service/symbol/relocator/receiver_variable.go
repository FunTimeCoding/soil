package relocator

import (
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func receiverVariable(
	p *packages.Package,
	declaration *ast.FuncDecl,
) *types.Var {
	names := declaration.Recv.List[0].Names

	if len(names) == 0 {
		return nil
	}

	v, _ := p.TypesInfo.Defs[names[0]].(*types.Var)

	return v
}
