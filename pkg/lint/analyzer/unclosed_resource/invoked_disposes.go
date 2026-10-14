package unclosed_resource

import (
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func (s *Summaries) invokedDisposes(
	p *packages.Package,
	l *ast.FuncLit,
	c *ast.CallExpr,
	o *types.Var,
) bool {
	var parameters []*types.Var

	for _, field := range l.Type.Params.List {
		for _, name := range field.Names {
			v, _ := p.TypesInfo.Defs[name].(*types.Var)
			parameters = append(parameters, v)
		}
	}

	for i, a := range c.Args {
		if i < len(parameters) && parameters[i] != nil &&
			rootIdentifier(p, a) == o && closes(p, l.Body, parameters[i], s) {
			return true
		}
	}

	return false
}
