package file

import "go/ast"

func findFunctionDeclaration(
	file *ast.File,
	name string,
) (*ast.FuncDecl, int) {
	for i, d := range file.Decls {
		f, okay := d.(*ast.FuncDecl)

		if okay && f.Name.Name == name {
			return f, i
		}
	}

	return nil, -1
}
