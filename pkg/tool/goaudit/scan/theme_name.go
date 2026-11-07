package scan

import (
	"github.com/funtimecoding/soil/pkg/tool/goaudit/constant"
	"go/ast"
	"strings"
)

func themeName(e ast.Expr) string {
	s, okay := e.(*ast.SelectorExpr)

	if !okay {
		return constant.UnknownTheme
	}

	return strings.TrimPrefix(s.Sel.Name, constant.ThemePrefix)
}
