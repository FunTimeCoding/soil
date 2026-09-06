package scan

import (
	"github.com/funtimecoding/soil/pkg/parse"
	"go/ast"
	"go/token"
)

func stringConstants(f *ast.File) []string {
	var result []string

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

			for i := range value.Names {
				if i >= len(value.Values) {
					continue
				}

				if literal, okay := parse.StringValue(value.Values[i]); okay {
					result = append(result, literal)
				}
			}
		}
	}

	return result
}
