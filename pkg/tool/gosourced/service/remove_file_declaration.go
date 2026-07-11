package service

import "go/ast"

func removeFileDeclaration(
	file *ast.File,
	declaration ast.Decl,
) {
	for i, d := range file.Decls {
		if d == declaration {
			file.Decls = append(file.Decls[:i], file.Decls[i+1:]...)

			return
		}
	}
}
