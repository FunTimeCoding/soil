package unclosed_resource

import (
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func assignedOnce(
	p *packages.Package,
	body *ast.BlockStmt,
	o *types.Var,
) bool {
	count := 0
	looped := false
	var walk func(n ast.Node, loop bool)
	walk = func(n ast.Node, loop bool) {
		ast.Inspect(
			n,
			func(c ast.Node) bool {
				switch v := c.(type) {
				case *ast.ForStmt:
					if v != n {
						walk(v, true)

						return false
					}
				case *ast.RangeStmt:
					if v != n {
						walk(v, true)

						return false
					}
				case *ast.AssignStmt:
					for _, l := range v.Lhs {
						i, okay := l.(*ast.Ident)

						if okay && (p.TypesInfo.Defs[i] == o || p.TypesInfo.Uses[i] == o) {
							count++
							looped = looped || loop
						}
					}
				}

				return true
			},
		)
	}
	walk(body, false)

	return count == 1 && !looped
}
