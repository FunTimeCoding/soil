package scan

import (
	"github.com/funtimecoding/soil/pkg/tool/goaudit/constant"
	"go/ast"
	"strings"
)

func themeName(c *ast.CallExpr) string {
	if len(c.Args) == 0 {
		return constant.UnknownTheme
	}

	s, okay := c.Args[0].(*ast.SelectorExpr)

	if !okay {
		return constant.UnknownTheme
	}

	return strings.TrimPrefix(s.Sel.Name, constant.ThemePrefix)
}
