package receiver_plan

import "go/ast"

type Plan struct {
	Method *ast.FuncDecl
	Idents []*ast.Ident
	Label  string
	Old    string
	New    string
}
