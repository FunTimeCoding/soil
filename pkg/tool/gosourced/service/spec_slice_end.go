package service

import (
	"go/ast"
	"go/token"
)

func specSliceEnd(spec ast.Spec) token.Pos {
	switch s := spec.(type) {
	case *ast.ValueSpec:
		if s.Comment != nil {
			return s.Comment.End()
		}
	case *ast.TypeSpec:
		if s.Comment != nil {
			return s.Comment.End()
		}
	}

	return spec.End()
}
