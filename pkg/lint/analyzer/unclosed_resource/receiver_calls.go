package unclosed_resource

import (
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func receiverCalls(
	p *packages.Package,
	f *ast.FuncDecl,
) []string {
	r := receiverOf(p, f)
	var result []string
	ast.Inspect(
		f.Body,
		func(n ast.Node) bool {
			c, okay := n.(*ast.CallExpr)

			if !okay {
				return true
			}

			selector, okay := c.Fun.(*ast.SelectorExpr)

			if !okay || !isReceiver(p, selector.X, r) {
				return true
			}

			s := p.TypesInfo.Selections[selector]

			if s != nil && s.Kind() == types.MethodVal && len(s.Index()) == 1 {
				result = append(result, selector.Sel.Name)
			}

			return true
		},
	)

	return result
}
