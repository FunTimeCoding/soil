package service

import "go/ast"

func declarationDocument(declaration ast.Decl) *ast.CommentGroup {
	switch d := declaration.(type) {
	case *ast.FuncDecl:
		return d.Doc
	case *ast.GenDecl:
		return d.Doc
	}

	return nil
}
