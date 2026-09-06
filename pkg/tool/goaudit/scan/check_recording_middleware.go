package scan

import (
	"github.com/funtimecoding/soil/pkg/system/virtual_file_system"
	"github.com/funtimecoding/soil/pkg/tool/goaudit/constant"
	"go/ast"
	"path/filepath"
)

func (s *Service) checkRecordingMiddleware(
	v *virtual_file_system.System,
	path string,
) {
	for _, name := range []string{constant.MountFileName, "run.go"} {
		f := parseWebFile(v, filepath.Join(path, name))

		if f == nil {
			continue
		}

		flagged := false
		ast.Inspect(
			f,
			func(n ast.Node) bool {
				if flagged {
					return false
				}

				c, okay := n.(*ast.CallExpr)

				if !okay {
					return true
				}

				m, okay := c.Fun.(*ast.SelectorExpr)

				if !okay ||
					m.Sel.Name != "NewStrictHandler" ||
					len(c.Args) < 2 {
					return true
				}

				if !hasRecordingMiddleware(c.Args[1]) {
					s.addConcern(
						constant.StrictMiddlewareKey,
						constant.StrictMiddlewareText,
						path,
					)
					flagged = true

					return false
				}

				return true
			},
		)

		if flagged {
			return
		}
	}
}
