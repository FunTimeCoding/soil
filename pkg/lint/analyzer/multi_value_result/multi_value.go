package multi_value_result

import (
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func multiValue(p *packages.Package, call *ast.CallExpr) bool {
	tuple, okay := p.TypesInfo.TypeOf(call).(*types.Tuple)

	return okay && tuple.Len() >= 2
}
