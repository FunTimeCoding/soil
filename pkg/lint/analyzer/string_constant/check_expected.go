package string_constant

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
	"strings"
)

// checkExpected flags constant references inside the expected
// argument of an assert call - expected values pin actuals and
// stay literal, a constant there drifts with the code under test.
func checkExpected(
	p *packages.Package,
	results *output.Results,
	expected ast.Expr,
) {
	ast.Inspect(
		expected,
		func(n ast.Node) bool {
			i, okay := n.(*ast.Ident)

			if !okay {
				return true
			}

			c, okay := p.TypesInfo.Uses[i].(*types.Const)

			if !okay || c.Pkg() == nil {
				return true
			}

			// Standard library constants cannot drift with the
			// code under test - they are named literals.
			if !strings.ContainsRune(
				strings.Split(c.Pkg().Path(), "/")[0],
				'.',
			) {
				return true
			}

			results.AddConcern(
				concern.NewFile(
					"string_constant",
					fmt.Sprintf(
						"constant %s.%s in expected value should be a literal",
						c.Pkg().Name(),
						c.Name(),
					),
					p.Fset.Position(i.Pos()).Filename,
					false,
				),
			)

			return true
		},
	)
}
