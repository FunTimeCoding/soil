package tester_function

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"go/ast"
)

func takesTestingFirst(f *ast.FuncDecl) bool {
	if f.Type.Params == nil || len(f.Type.Params.List) == 0 {
		return false
	}

	v, okay := f.Type.Params.List[0].Type.(*ast.StarExpr)

	if !okay {
		return false
	}

	s, okay := v.X.(*ast.SelectorExpr)

	if !okay {
		return false
	}

	i, okay := s.X.(*ast.Ident)

	if !okay {
		return false
	}

	return join.Empty(
		"*",
		i.Name,
		constant.MemberSeparator,
		s.Sel.Name,
	) == constant.AssertTestingType
}
