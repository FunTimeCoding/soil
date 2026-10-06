package service

import (
	"go/ast"
	"go/token"
)

func declaringStatement(path []ast.Node) (ast.Stmt, *ast.BlockStmt, ast.Expr) {
	if len(path) < 3 {
		return nil, nil, nil
	}

	switch s := path[1].(type) {
	case *ast.AssignStmt:
		block, okay := path[2].(*ast.BlockStmt)

		if !okay || s.Tok != token.DEFINE || len(s.Lhs) != 1 || len(s.Rhs) != 1 {
			return nil, nil, nil
		}

		return s, block, s.Rhs[0]
	case *ast.ValueSpec:
		if len(path) < 5 || len(s.Names) != 1 {
			return nil, nil, nil
		}

		g, isGeneral := path[2].(*ast.GenDecl)
		statement, isDeclaration := path[3].(*ast.DeclStmt)
		block, isBlock := path[4].(*ast.BlockStmt)

		if !isGeneral || !isDeclaration || !isBlock || len(g.Specs) != 1 {
			return nil, nil, nil
		}

		if len(s.Values) == 1 {
			return statement, block, s.Values[0]
		}

		return statement, block, nil
	}

	return nil, nil, nil
}
