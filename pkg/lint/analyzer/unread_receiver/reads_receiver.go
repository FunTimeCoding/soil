package unread_receiver

import (
	"go/ast"
	"golang.org/x/tools/go/packages"
)

func readsReceiver(
	p *packages.Package,
	f *ast.FuncDecl,
) bool {
	names := f.Recv.List[0].Names

	if len(names) == 0 {
		return false
	}

	o := p.TypesInfo.Defs[names[0]]
	result := false
	ast.Inspect(
		f.Body,
		func(n ast.Node) bool {
			if i, okay := n.(*ast.Ident); okay && p.TypesInfo.Uses[i] == o {
				result = true
			}

			return !result
		},
	)

	return result
}
