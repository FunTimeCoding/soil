package service

import (
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/literal"
	"go/ast"
)

func assignmentTarget(site *literal.Site) (ast.Stmt, ast.Expr) {
	var statement ast.Stmt
	var target ast.Expr

	switch p := site.Parent.(type) {
	case *ast.AssignStmt:
		statement, target = p, p.Lhs[0]
	case *ast.ValueSpec:
		for _, n := range site.Path {
			if d, okay := n.(*ast.DeclStmt); okay {
				statement, target = d, p.Names[0]

				break
			}
		}
	}

	if statement == nil {
		return nil, nil
	}

	if i, okay := target.(*ast.Ident); okay && i.Name == "_" {
		return nil, nil
	}

	for i, n := range site.Path {
		if n != statement || i+1 >= len(site.Path) {
			continue
		}

		switch site.Path[i+1].(type) {
		case *ast.BlockStmt, *ast.CaseClause, *ast.CommClause:
			return statement, target
		}
	}

	return nil, nil
}
