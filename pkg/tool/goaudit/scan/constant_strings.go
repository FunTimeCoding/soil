package scan

import (
	library "github.com/funtimecoding/soil/pkg/constant"
	"github.com/funtimecoding/soil/pkg/parse"
	"github.com/funtimecoding/soil/pkg/system/virtual_file_system"
	"github.com/funtimecoding/soil/pkg/tool/goaudit/constant"
	"go/ast"
	"go/token"
	"path/filepath"
	"strings"
)

func constantStrings(v *virtual_file_system.System) map[string]string {
	result := make(map[string]string)

	if !v.DirectoryExists(constant.ConstantDirectory) {
		return result
	}

	for _, name := range v.MustReadDirectory(constant.ConstantDirectory) {
		if !strings.HasSuffix(name, library.GoExtension) {
			continue
		}

		file := filepath.Join(constant.ConstantDirectory, name)
		f, _, e := parse.Source(name, v.ReadString(file))

		if e != nil {
			continue
		}

		for _, d := range f.Decls {
			g, okay := d.(*ast.GenDecl)

			if !okay || g.Tok != token.CONST {
				continue
			}

			for _, s := range g.Specs {
				value, okay := s.(*ast.ValueSpec)

				if !okay {
					continue
				}

				for i, n := range value.Names {
					if i >= len(value.Values) {
						continue
					}

					if literal, okay := parse.StringValue(
						value.Values[i],
					); okay {
						result[n.Name] = literal
					}
				}
			}
		}
	}

	return result
}
