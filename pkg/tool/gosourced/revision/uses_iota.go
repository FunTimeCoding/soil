package revision

import (
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"go/ast"
)

func usesIota(d *ast.GenDecl) bool {
	result := false
	ast.Inspect(
		d,
		func(c ast.Node) bool {
			if i, isIdent := c.(*ast.Ident); isIdent &&
				i.Name == constant.IotaName {
				result = true
			}

			return !result
		},
	)

	return result
}
