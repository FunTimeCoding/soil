package gofix

import (
	"go/ast"
	"go/types"
)

type Violation struct {
	ident   *ast.Ident
	object  types.Object
	segment string
	fix     string
}
