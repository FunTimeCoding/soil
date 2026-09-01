package scan

import (
	libraryConstant "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/parse"
	"github.com/funtimecoding/soil/pkg/system/virtual_file_system"
	"github.com/funtimecoding/soil/pkg/tool/goaudit/constant"
	"go/ast"
	"path/filepath"
)

func (s *Service) checkWebArgument(
	v *virtual_file_system.System,
	path string,
) {
	file := filepath.Join(path, libraryConstant.MainFile)

	if !v.Has(file) {
		return
	}

	f, _, e := parse.Source(libraryConstant.MainFile, v.ReadString(file))

	if e != nil {
		return
	}

	ast.Inspect(
		f,
		func(n ast.Node) bool {
			c, okay := n.(*ast.CallExpr)

			if !okay || len(c.Args) == 0 {
				return true
			}

			m, okay := c.Fun.(*ast.SelectorExpr)

			if !okay || m.Sel.Name != "Integer" {
				return true
			}

			a, okay := c.Args[0].(*ast.SelectorExpr)

			if !okay {
				return true
			}

			if a.Sel.Name == "Port" || a.Sel.Name == "MetricPort" {
				s.addConcern(
					constant.WebArgumentKey,
					constant.WebArgumentText,
					path,
				)

				return false
			}

			return true
		},
	)
}
