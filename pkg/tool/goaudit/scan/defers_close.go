package scan

import (
	"github.com/funtimecoding/soil/pkg/tool/goaudit/constant"
	"go/ast"
)

func defersClose(f *ast.File) bool {
	found := false
	ast.Inspect(
		f,
		func(n ast.Node) bool {
			if found {
				return false
			}

			d, okay := n.(*ast.DeferStmt)

			if !okay {
				return true
			}

			if m, okay := d.Call.Fun.(*ast.SelectorExpr); okay &&
				m.Sel.Name == constant.CloseMethod {
				found = true
			}

			return true
		},
	)

	return found
}
