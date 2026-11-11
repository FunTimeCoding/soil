package symbol

import (
	"fmt"
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func selectionShape(
	p *packages.Package,
	method string,
	receiverName string,
	s *ast.SelectorExpr,
	target *types.Named,
	moved map[string]bool,
) string {
	selection := p.TypesInfo.Selections[s]

	if selection == nil {
		return ""
	}

	name := s.Sel.Name
	found, _, _ := types.LookupFieldOrMethod(
		types.NewPointer(target),
		true,
		target.Obj().Pkg(),
		name,
	)

	switch selection.Kind() {
	case types.FieldVal:
		field, okay := found.(*types.Var)

		if !okay {
			return fmt.Sprintf(
				"%s reads %s.%s, which %s lacks",
				method,
				receiverName,
				name,
				qualifiedType(target),
			)
		}

		if !types.Identical(field.Type(), selection.Obj().Type()) {
			return fmt.Sprintf(
				"%s reads %s.%s as %s, which %s has as %s",
				method,
				receiverName,
				name,
				selection.Obj().Type(),
				qualifiedType(target),
				field.Type(),
			)
		}
	case types.MethodVal:
		if moved[name] {
			return ""
		}

		f, okay := found.(*types.Func)

		if !okay || !types.Identical(f.Type(), selection.Obj().Type()) {
			return fmt.Sprintf(
				"%s calls %s.%s, which is neither moved nor declared on %s",
				method,
				receiverName,
				name,
				qualifiedType(target),
			)
		}
	}

	return ""
}
