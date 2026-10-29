package scan

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goaudit/constant"
	"go/ast"
	"slices"
	"strings"
)

func runConstructors(f *ast.File) []string {
	var result []string

	for _, d := range f.Decls {
		g, okay := d.(*ast.FuncDecl)

		if !okay || g.Name.Name != "Run" || g.Body == nil {
			continue
		}

		ast.Inspect(
			g.Body,
			func(n ast.Node) bool {
				c, okay := n.(*ast.CallExpr)

				if !okay {
					return true
				}

				m, okay := c.Fun.(*ast.SelectorExpr)

				if !okay {
					return true
				}

				i, okay := m.X.(*ast.Ident)

				if !okay || slices.Contains(constant.RunPlumbing, i.Name) {
					return true
				}

				if !strings.HasPrefix(m.Sel.Name, "New") &&
					!strings.HasPrefix(m.Sel.Name, constant.OpenPrefix) {
					return true
				}

				name := fmt.Sprintf("%s.%s", i.Name, m.Sel.Name)

				if !slices.Contains(result, name) {
					result = append(result, name)
				}

				return true
			},
		)
	}

	return result
}
