package revision

import (
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"go/ast"
)

func declaredNames(f *ast.File) []string {
	var result []string

	for _, d := range f.Decls {
		switch d := d.(type) {
		case *ast.FuncDecl:
			if d.Recv == nil {
				result = append(result, d.Name.Name)
			}
		case *ast.GenDecl:
			for _, spec := range d.Specs {
				result = append(result, specNames(spec)...)
			}
		}
	}

	var named []string

	for _, name := range result {
		if name != constant.BlankName {
			named = append(named, name)
		}
	}

	return named
}
