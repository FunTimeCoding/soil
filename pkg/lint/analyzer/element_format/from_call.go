package element_format

import "go/ast"

func FromCall(call *ast.CallExpr) *Elements {
	return &Elements{
		Open:     call.Lparen,
		Close:    call.Rparen,
		Items:    call.Args,
		Position: call.Pos(),
		Ellipsis: call.Ellipsis.IsValid(),
	}
}
