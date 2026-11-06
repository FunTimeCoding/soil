package stub_constructor

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"go/ast"
	"golang.org/x/tools/go/packages"
	"strings"
)

func gormModels(p *packages.Package) map[string]bool {
	result := map[string]bool{}

	for _, file := range p.Syntax {
		ast.Inspect(
			file,
			func(n ast.Node) bool {
				t, okay := n.(*ast.TypeSpec)

				if !okay {
					return true
				}

				s, okay := t.Type.(*ast.StructType)

				if !okay {
					return true
				}

				for _, field := range s.Fields.List {
					if field.Tag != nil &&
						strings.Contains(field.Tag.Value, constant.GormTag) {
						result[t.Name.Name] = true
					}
				}

				return true
			},
		)
	}

	return result
}
