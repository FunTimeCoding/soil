package nil_dependency

import (
	"go/ast"
	"go/types"
)

func parameterType(
	s *types.Signature,
	call *ast.CallExpr,
	index int,
) types.Type {
	last := s.Params().Len() - 1

	if !s.Variadic() || index < last {
		if index > last {
			return nil
		}

		return s.Params().At(index).Type()
	}

	if call.Ellipsis.IsValid() {
		return s.Params().At(last).Type()
	}

	if slice, okay := s.Params().At(last).Type().(*types.Slice); okay {
		return slice.Elem()
	}

	return nil
}
