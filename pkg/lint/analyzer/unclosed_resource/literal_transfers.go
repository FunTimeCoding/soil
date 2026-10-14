package unclosed_resource

import (
	"github.com/funtimecoding/soil/pkg/lint/types/resource_transfer"
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func literalTransfers(
	p *packages.Package,
	l *ast.CompositeLit,
) []resource_transfer.Transfer {
	named, okay := dereference(p.TypesInfo.TypeOf(l)).(*types.Named)

	if !okay {
		return nil
	}

	s, okay := named.Underlying().(*types.Struct)

	if !okay {
		return nil
	}

	var result []resource_transfer.Transfer

	for i, e := range l.Elts {
		var field *types.Var
		value := e

		if pair, isPair := e.(*ast.KeyValueExpr); isPair {
			key, isIdent := pair.Key.(*ast.Ident)

			if !isIdent {
				continue
			}

			field, _ = p.TypesInfo.Uses[key].(*types.Var)
			value = pair.Value
		} else if i < s.NumFields() {
			field = s.Field(i)
		}

		if field == nil {
			continue
		}

		result = append(result, fieldTransfers(named, field, value)...)
	}

	return result
}
