package call_format

import (
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"go/ast"
	"golang.org/x/tools/go/packages"
)

func reportMultiLine(
	p *packages.Package,
	results *output.Results,
	call *ast.CallExpr,
) {
	openParenLine := p.Fset.Position(call.Lparen).Line
	firstArgLine := p.Fset.Position(call.Args[0].Pos()).Line

	if openParenLine == firstArgLine {
		results.AddConcern(
			concern.NewFile(
				"call_format",
				"each argument should be on its own line",
				p.Fset.Position(call.Args[0].Pos()).Filename,
				false,
			),
		)

		return
	}

	for i := 1; i < len(call.Args); i++ {
		previousEnd := p.Fset.Position(call.Args[i-1].End()).Line
		currentStart := p.Fset.Position(call.Args[i].Pos()).Line

		if previousEnd == currentStart {
			results.AddConcern(
				concern.NewFile(
					"call_format",
					"each argument should be on its own line",
					p.Fset.Position(call.Args[i].Pos()).Filename,
					false,
				),
			)

			return
		}
	}
}
