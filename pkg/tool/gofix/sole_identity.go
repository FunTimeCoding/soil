package gofix

import "go/ast"

func soleIdentity(f *ast.File) string {
	var names []string

	for _, d := range f.Decls {
		switch v := d.(type) {
		case *ast.GenDecl:
			for _, s := range v.Specs {
				if t, okay := s.(*ast.TypeSpec); okay {
					names = append(names, t.Name.Name)
				}
			}
		case *ast.FuncDecl:
			names = append(names, v.Name.Name)
		}
	}

	if len(names) != 1 {
		return ""
	}

	return names[0]
}
