package pattern

import (
	"github.com/dave/dst"
	"github.com/dave/dst/dstutil"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/decoration"
	"go/ast"
	"golang.org/x/tools/go/packages"
)

func substitute(
	node dst.Node,
	holes map[string]string,
	querySymbol string,
	qualifiers map[string]string,
	set *decoration.Set,
	p *packages.Package,
	bindings map[string]ast.Expr,
	anchor ast.Node,
) dst.Node {
	return dstutil.Apply(
		node,
		func(c *dstutil.Cursor) bool {
			switch n := c.Node().(type) {
			case *dst.CallExpr:
				if callReferences(n, querySymbol) && spreadHoleCall(n, holes) {
					twin, okay := set.DecoratedNode(p, anchor).(*dst.CallExpr)

					if okay {
						call := dst.Clone(n).(*dst.CallExpr)
						call.Fun = substitute(
							call.Fun,
							holes,
							querySymbol,
							qualifiers,
							set,
							p,
							bindings,
							anchor,
						).(dst.Expr)
						call.Args = dst.Clone(twin).(*dst.CallExpr).Args
						call.Ellipsis = twin.Ellipsis
						c.Replace(call)
					}

					return false
				}
			case *dst.SelectorExpr:
				qualifier, plain := n.X.(*dst.Ident)

				if plain {
					if path, known := qualifiers[qualifier.Name]; known {
						c.Replace(&dst.Ident{Name: n.Sel.Name, Path: path})

						return false
					}
				}
			case *dst.Ident:
				if _, isHole := holes[n.Name]; isHole {
					twin := set.DecoratedNode(p, bindings[n.Name])

					if twin != nil {
						c.Replace(dst.Clone(twin))
					}

					return false
				}
			}

			return true
		},
		nil,
	)
}
