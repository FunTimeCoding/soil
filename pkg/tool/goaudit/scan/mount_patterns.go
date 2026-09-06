package scan

import (
	"github.com/funtimecoding/soil/pkg/system/virtual_file_system"
	"github.com/funtimecoding/soil/pkg/tool/goaudit/constant"
	"go/ast"
	"path/filepath"
)

func mountPatterns(
	v *virtual_file_system.System,
	path string,
	constants map[string]string,
) []string {
	var result []string

	for _, name := range []string{constant.MountFileName, "run.go"} {
		f := parseWebFile(v, filepath.Join(path, name))

		if f == nil {
			continue
		}

		ast.Inspect(
			f,
			func(n ast.Node) bool {
				c, okay := n.(*ast.CallExpr)

				if !okay || len(c.Args) < 2 {
					return true
				}

				m, okay := c.Fun.(*ast.SelectorExpr)

				if !okay {
					return true
				}

				switch m.Sel.Name {
				case "TokenMount", "OpenMount", "Token", "Open", "Session",
					"SessionMount":
				default:
					return true
				}

				if pattern, okay := evaluatePattern(
					c.Args[0],
					constants,
				); okay {
					result = append(result, pattern)
				}

				return true
			},
		)
	}

	return result
}
