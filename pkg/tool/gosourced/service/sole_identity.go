package service

import (
	"go/ast"
	"go/token"
)

func soleIdentity(file *ast.File) (string, bool) {
	var names []string

	for _, d := range file.Decls {
		switch t := d.(type) {
		case *ast.FuncDecl:
			names = append(names, t.Name.Name)
		case *ast.GenDecl:
			if t.Tok != token.TYPE {
				continue
			}

			for _, s := range t.Specs {
				if v, isType := s.(*ast.TypeSpec); isType {
					names = append(names, v.Name.Name)
				}
			}
		}
	}

	if len(names) != 1 {
		return "", false
	}

	return names[0], true
}
