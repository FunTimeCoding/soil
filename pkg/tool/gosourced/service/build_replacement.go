package service

import (
	"fmt"
	"github.com/dave/dst"
	"github.com/dave/dst/decorator"
	"github.com/dave/dst/dstutil"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/decoration"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/match"
	"go/ast"
	"golang.org/x/tools/go/packages"
)

func buildReplacement(
	replacement *match.Pattern,
	source string,
	querySymbol string,
	qualifiers map[string]string,
	set *decoration.Set,
	p *packages.Package,
	bindings map[string]ast.Expr,
	anchor ast.Node,
) (dst.Stmt, error) {
	file, e := decorator.Parse(join.Empty("package replacement\n\n", source))

	if e != nil {
		return nil, e
	}

	var statement dst.Stmt

	for _, d := range file.Decls {
		declaration, okay := d.(*dst.FuncDecl)

		if okay && len(declaration.Body.List) == 1 {
			statement = declaration.Body.List[0]
		}
	}

	if statement == nil {
		return nil, fmt.Errorf("replacement holds no statement")
	}

	result := dstutil.Apply(
		statement,
		func(c *dstutil.Cursor) bool {
			switch n := c.Node().(type) {
			case *dst.CallExpr:
				if callReferences(n, querySymbol) &&
					spreadHoleCall(n, replacement.Holes) {
					twin := set.DecoratedNode(p, anchor)

					if twin != nil {
						c.Replace(dst.Clone(twin))
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
				if _, isHole := replacement.Holes[n.Name]; isHole {
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

	return result.(dst.Stmt), nil
}
