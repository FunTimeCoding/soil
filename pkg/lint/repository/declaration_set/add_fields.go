package declaration_set

import "go/ast"

func (s *Set) addFields(
	owner string,
	fields *ast.FieldList,
) {
	if fields == nil {
		return
	}

	for _, f := range fields.List {
		for _, n := range f.Names {
			s.addMember(owner, n.Name)
		}

		if len(f.Names) > 0 {
			continue
		}

		switch embedded := typeExpression(f.Type).(type) {
		case *ast.Ident:
			s.addMember(owner, embedded.Name)
			s.embeds[owner] = append(s.embeds[owner], embedded.Name)
		case *ast.SelectorExpr:
			s.addMember(owner, embedded.Sel.Name)
			s.foreign[owner] = true
		}
	}
}
