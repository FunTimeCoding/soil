package revision

import "go/ast"

func declaredInStatement(statement ast.Stmt) []string {
	d, isDeclaration := statement.(*ast.DeclStmt)

	if !isDeclaration {
		return nil
	}

	g, isGeneral := d.Decl.(*ast.GenDecl)

	if !isGeneral {
		return nil
	}

	var result []string

	for _, spec := range g.Specs {
		result = append(result, specNames(spec)...)
	}

	return result
}
