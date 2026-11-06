package stub_constructor

import (
	"go/ast"
	"go/token"
)

func emptyReturn(f *ast.FuncDecl) string {
	if f.Type.Params.NumFields() > 0 ||
		f.Type.Results.NumFields() != 1 ||
		f.Body == nil ||
		len(f.Body.List) != 1 {
		return ""
	}

	pointer, okay := f.Type.Results.List[0].Type.(*ast.StarExpr)

	if !okay {
		return ""
	}

	result, okay := pointer.X.(*ast.Ident)

	if !okay {
		return ""
	}

	r, okay := f.Body.List[0].(*ast.ReturnStmt)

	if !okay || len(r.Results) != 1 {
		return ""
	}

	u, okay := r.Results[0].(*ast.UnaryExpr)

	if !okay || u.Op != token.AND {
		return ""
	}

	c, okay := u.X.(*ast.CompositeLit)

	if !okay || len(c.Elts) > 0 {
		return ""
	}

	literal, okay := c.Type.(*ast.Ident)

	if !okay || literal.Name != result.Name {
		return ""
	}

	return result.Name
}
