package receiver_plan

import "go/ast"

func New(
	method *ast.FuncDecl,
	idents []*ast.Ident,
	label string,
	old string,
	newValue string,
) *Plan {
	return &Plan{
		Method: method,
		Idents: idents,
		Label:  label,
		Old:    old,
		New:    newValue,
	}
}
