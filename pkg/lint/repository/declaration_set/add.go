package declaration_set

import "go/ast"

func (s *Set) Add(d ast.Decl) {
	switch declaration := d.(type) {
	case *ast.FuncDecl:
		if declaration.Recv == nil {
			s.Names[declaration.Name.Name] = true

			return
		}

		if owner, okay := typeExpression(
			declaration.Recv.List[0].Type,
		).(*ast.Ident); okay {
			s.addMember(owner.Name, declaration.Name.Name)
		}
	case *ast.GenDecl:
		for _, spec := range declaration.Specs {
			switch v := spec.(type) {
			case *ast.TypeSpec:
				s.Names[v.Name.Name] = true
				s.Types[v.Name.Name] = true

				switch t := v.Type.(type) {
				case *ast.StructType:
					s.addFields(v.Name.Name, t.Fields)
				case *ast.InterfaceType:
					s.addFields(v.Name.Name, t.Methods)
				}
			case *ast.ValueSpec:
				for _, n := range v.Names {
					s.Names[n.Name] = true
				}
			}
		}
	}
}
