package scan

import (
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/system/virtual_file_system"
	"github.com/funtimecoding/soil/pkg/tool/goaudit/constant"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
	"go/ast"
	"go/token"
	"path/filepath"
	"strconv"
)

func pathConstants(
	v *virtual_file_system.System,
	path string,
) map[string]string {
	result := map[string]string{
		"Slash":            stringsConstant.Slash,
		"RootPath":         webConstant.RootPath,
		"RootPattern":      webConstant.RootPattern,
		"InterfacePath":    webConstant.InterfacePath,
		"LivePath":         webConstant.LivePath,
		"SearchPath":       webConstant.SearchPath,
		"FaviconPath":      webConstant.FaviconPath,
		"MetricsPath":      webConstant.MetricsPath,
		"ModelContextPath": webConstant.ModelContextPath,
		"EventPath":        webConstant.EventPath,
		"SignInPath":       webConstant.SignInPath,
		"CallbackPath":     webConstant.CallbackPath,
		"SignOutPath":      webConstant.SignOutPath,
	}
	directory := filepath.Join(path, constant.ConstantDirectory)

	if !v.DirectoryExists(directory) {
		return result
	}

	for _, name := range v.MustReadDirectory(directory) {
		f := parseWebFile(v, filepath.Join(directory, name))

		if f == nil {
			continue
		}

		ast.Inspect(
			f,
			func(n ast.Node) bool {
				d, okay := n.(*ast.GenDecl)

				if !okay || d.Tok != token.CONST {
					return true
				}

				for _, spec := range d.Specs {
					s, okay := spec.(*ast.ValueSpec)

					if !okay || len(s.Names) != len(s.Values) {
						continue
					}

					for i, identifier := range s.Names {
						l, okay := s.Values[i].(*ast.BasicLit)

						if !okay || l.Kind != token.STRING {
							continue
						}

						value, e := strconv.Unquote(l.Value)

						if e == nil {
							result[identifier.Name] = value
						}
					}
				}

				return false
			},
		)
	}

	return result
}
