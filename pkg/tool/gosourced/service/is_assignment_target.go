package service

import (
	"go/ast"
	"go/token"
)

func isAssignmentTarget(path []ast.Node) bool {
	if len(path) < 2 {
		return false
	}

	target := path[0]
	k := 1

	if s, okay := path[1].(*ast.SelectorExpr); okay &&
		s.Sel.Pos() == target.Pos() {
		target = s
		k = 2
	}

	if k >= len(path) {
		return false
	}

	switch s := path[k].(type) {
	case *ast.AssignStmt:
		if s.Tok == token.DEFINE {
			return false
		}

		for _, l := range s.Lhs {
			if l.Pos() == target.Pos() && l.End() == target.End() {
				return true
			}
		}
	case *ast.IncDecStmt:
		return s.X.Pos() == target.Pos() && s.X.End() == target.End()
	case *ast.KeyValueExpr:
		return s.Key.Pos() == target.Pos()
	}

	return false
}
