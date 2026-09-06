package scan

import (
	"github.com/funtimecoding/soil/pkg/system/virtual_file_system"
	"github.com/funtimecoding/soil/pkg/tool/goaudit/constant"
	"go/ast"
	"path/filepath"
)

func (s *Service) checkModelContext(
	v *virtual_file_system.System,
	path string,
) {
	if !s.ModelContext {
		return
	}

	mc := filepath.Join(path, constant.ModelContextDirectory)

	if !v.Has(filepath.Join(mc, constant.MountFileName)) {
		s.addConcern(constant.MissingMountKey, constant.MissingMountText, path)
	}

	if !v.Has(filepath.Join(mc, "capture_fail.go")) {
		s.addConcern(
			constant.MissingCaptureFailKey,
			constant.MissingCaptureFailText,
			path,
		)
	}

	if v.Has(filepath.Join(mc, "nested.go")) {
		s.addConcern(constant.StaleNestedKey, constant.StaleNestedText, path)
	}

	f := parseWebFile(v, filepath.Join(mc, constant.NewFileName))

	if f == nil {
		return
	}

	ast.Inspect(
		f,
		func(n ast.Node) bool {
			c, okay := n.(*ast.CallExpr)

			if !okay {
				return true
			}

			if m, okay := c.Fun.(*ast.SelectorExpr); okay &&
				m.Sel.Name == "NewMCPServer" {
				s.addConcern(
					constant.RawModelContextServerKey,
					constant.RawModelContextServerText,
					path,
				)

				return false
			}

			return true
		},
	)
}
