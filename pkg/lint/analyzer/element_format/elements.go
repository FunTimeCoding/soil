package element_format

import (
	"go/ast"
	"go/token"
)

type Elements struct {
	Open        token.Pos
	Close       token.Pos
	Items       []ast.Expr
	Position    token.Pos
	Ellipsis    bool
	Padding     int
	HasComments bool
}
