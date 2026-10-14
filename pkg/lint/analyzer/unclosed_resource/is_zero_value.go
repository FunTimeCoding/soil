package unclosed_resource

import (
	"go/ast"
	"go/token"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func isZeroValue(
	p *packages.Package,
	body *ast.BlockStmt,
	e ast.Expr,
) bool {
	if isNil(p, e) {
		return true
	}

	i, okay := ast.Unparen(e).(*ast.Ident)

	if !okay {
		return false
	}

	o, okay := p.TypesInfo.Uses[i].(*types.Var)

	if !okay || len(assignedValues(p, body, o)) > 0 {
		return false
	}

	declared := false
	ast.Inspect(
		body,
		func(n ast.Node) bool {
			d, isDeclaration := n.(*ast.DeclStmt)

			if !isDeclaration {
				return !declared
			}

			g, isGeneric := d.Decl.(*ast.GenDecl)

			if !isGeneric || g.Tok != token.VAR {
				return true
			}

			for _, spec := range g.Specs {
				v, isValue := spec.(*ast.ValueSpec)

				if !isValue || len(v.Values) > 0 {
					continue
				}

				for _, name := range v.Names {
					if p.TypesInfo.Defs[name] == o {
						declared = true
					}
				}
			}

			return !declared
		},
	)

	return declared
}
