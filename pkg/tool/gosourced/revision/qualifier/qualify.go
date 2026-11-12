package qualifier

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"go/ast"
	"strconv"
)

func (q *Qualifier) Qualify(
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

				if p, imported := q.Imports[x.Name]; imported {
					x.Name = strconv.Quote(p)
					skip[x] = true
				}
			case *ast.Ident:
				if !skip[c] && !locals[c.Name] && q.Own[c.Name] {
					c.Name = join.Empty(
						strconv.Quote(q.Path),
						constant.MemberSeparator,
						c.Name,
					)
				}
			}

			return true
		},
	)
}
