package scan

import (
	"github.com/funtimecoding/soil/pkg/parse"
	"github.com/funtimecoding/soil/pkg/tool/goaudit/constant"
	"go/ast"
)

func mountsView(f *ast.File) bool {
	for _, c := range parse.FindMethods(f, constant.MountMethod) {
		s, okay := c.Fun.(*ast.SelectorExpr)

		if !okay {
			continue
		}

		if field, okay := s.X.(*ast.SelectorExpr); okay &&
			field.Sel.Name == constant.ViewField {
			return true
		}
	}

	return false
}
