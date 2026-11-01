package revision

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"go/ast"
	"strconv"
)

func (q *qualifier) qualify(
	n ast.Node,
	locals map[string]bool,
) {
	skip := skipped(n)
	ast.Inspect(
		n,
		func(c ast.Node) bool {
			switch c := c.(type) {
			case *ast.SelectorExpr:
				x, isIdent := c.X.(*ast.Ident)

				if !isIdent || locals[x.Name] {
					return true
				}

				if p, imported := q.imports[x.Name]; imported {
					x.Name = strconv.Quote(p)
					skip[x] = true
				}
			case *ast.Ident:
				if !skip[c] && !locals[c.Name] && q.own[c.Name] {
					c.Name = join.Empty(
						strconv.Quote(q.path),
						constant.MemberSeparator,
						c.Name,
					)
				}
			}

			return true
		},
	)
}
