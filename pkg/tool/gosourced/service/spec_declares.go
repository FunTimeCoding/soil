package service

import (
	"go/ast"
	"go/types"
)

func specDeclares(
	s ast.Spec,
	o types.Object,
) bool {
	switch spec := s.(type) {
	case *ast.ValueSpec:
		for _, n := range spec.Names {
			if n.Pos() == o.Pos() {
				return true
			}
		}
	case *ast.TypeSpec:
		return spec.Name.Pos() == o.Pos()
	}

	return false
}
