package unclosed_resource

import (
	"go/ast"
	"golang.org/x/tools/go/packages"
)

func returnsSelf(
	p *packages.Package,
	f *ast.FuncDecl,
) bool {
	r := receiverOf(p, f)

	if r == nil {
		return false
	}

	found := false
	self := true
	ast.Inspect(
		f.Body,
		func(n ast.Node) bool {
			switch v := n.(type) {
			case *ast.FuncLit:
				return false
			case *ast.ReturnStmt:
				found = true
				self = self && len(v.Results) == 1 &&
					isReceiver(p, v.Results[0], r)
			}

			return self
		},
	)

	return found && self
}
