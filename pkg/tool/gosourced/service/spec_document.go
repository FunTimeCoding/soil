package service

import "go/ast"

func specDocument(
	declaration ast.Decl,
	spec ast.Spec,
) *ast.CommentGroup {
	switch s := spec.(type) {
	case *ast.ValueSpec:
		if s.Doc != nil {
			return s.Doc
		}
	case *ast.TypeSpec:
		if s.Doc != nil {
			return s.Doc
		}
	}

	g, okay := declaration.(*ast.GenDecl)

	if okay && len(g.Specs) == 1 {
		return g.Doc
	}

	return nil
}
