package scan

import (
	"github.com/funtimecoding/soil/pkg/tool/goaudit/constant"
	"go/ast"
	"go/token"
)

func declaresService(f *ast.File) bool {
	for _, d := range f.Decls {
		g, okay := d.(*ast.GenDecl)

		if !okay || g.Tok != token.TYPE {
			continue
		}

		for _, p := range g.Specs {
			if t, okay := p.(*ast.TypeSpec); okay &&
				t.Name.Name == constant.ServiceType {
				return true
			}
		}
	}

	return false
}
