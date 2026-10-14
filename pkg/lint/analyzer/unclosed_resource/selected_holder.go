package unclosed_resource

import (
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func selectedHolder(
	p *packages.Package,
	s *ast.SelectorExpr,
) *types.Named {
	selection := p.TypesInfo.Selections[s]

	if selection == nil || selection.Kind() != types.FieldVal {
		return nil
	}

	t := selection.Recv()
	index := selection.Index()

	for _, i := range index[:len(index)-1] {
		structure, okay := dereference(t).Underlying().(*types.Struct)

		if !okay {
			return nil
		}

		t = structure.Field(i).Type()
	}

	result, _ := dereference(t).(*types.Named)

	return result
}
