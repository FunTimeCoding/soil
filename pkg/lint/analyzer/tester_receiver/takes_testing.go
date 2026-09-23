package tester_receiver

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"go/ast"
)

func takesTesting(f *ast.FuncDecl) bool {
	if f.Type.Params == nil {
		return false
	}

	for _, field := range f.Type.Params.List {
		buffer := ""

		switch v := field.Type.(type) {
		case *ast.StarExpr:
			s, okay := v.X.(*ast.SelectorExpr)

			if !okay {
				continue
			}

			i, okay := s.X.(*ast.Ident)

			if !okay {
				continue
			}

			buffer = join.Empty(
				"*",
				i.Name,
				constant.MemberSeparator,
				s.Sel.Name,
			)
		default:
			continue
		}

		if buffer == constant.AssertTestingType {
			return true
		}
	}

	return false
}
