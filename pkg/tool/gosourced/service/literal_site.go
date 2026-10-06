package service

import (
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/literal"
	"go/ast"
	"go/token"
	"go/types"
	"golang.org/x/tools/go/ast/astutil"
	"golang.org/x/tools/go/packages"
)

func literalSite(
	p *packages.Package,
	file *ast.File,
	structure *types.Struct,
	target ast.Expr,
) *literal.Site {
	path, _ := astutil.PathEnclosingInterval(file, target.Pos(), target.End())
	index := 0

	for index < len(path)-1 && path[index] != target {
		index++
	}

	outer := target

	if u, isUnary := path[index+1].(*ast.UnaryExpr); isUnary &&
		u.Op == token.AND {
		outer = u
		index++
	}

	result := literal.NewSite(p, file, path, target, outer, path[index+1])

	if lit, okay := target.(*ast.CompositeLit); okay {
		result.Fields, _ = literalFields(structure, lit)
	}

	result.Nested = !standalone(result.Parent, outer)
	result.Shape = literalShape(structure, target, outer, result.Nested)

	return result
}
