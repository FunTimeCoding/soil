package naming

import (
	"github.com/funtimecoding/soil/pkg/lint/output"
	"go/ast"
	"golang.org/x/tools/go/packages"
)

func Check(
	p *packages.Package,
	results *output.Results,
) {
	for _, file := range p.Syntax {
		if ast.IsGenerated(file) {
			continue
		}

		ast.Inspect(
			file,
			func(n ast.Node) bool {
				ident, okay := n.(*ast.Ident)

				if !okay {
					return true
				}

				o, isDefined := p.TypesInfo.Defs[ident]

				if !isDefined {
					return true
				}

				checkIdent(p, results, ident, o)

				return true
			},
		)
	}
}
