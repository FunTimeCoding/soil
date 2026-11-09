package symbol

import "go/ast"

type plannedReceiver struct {
	Method *ast.FuncDecl
	Idents []*ast.Ident
	Label  string
	Old    string
	New    string
}
