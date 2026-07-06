package call_format

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/suppress"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"go/ast"
	"golang.org/x/tools/go/packages"
)

func checkCall(
	p *packages.Package,
	results *output.Results,
	call *ast.CallExpr,
) {
	if len(call.Args) < 2 {
		return
	}

	if suppress.IsSuppressed(p.Fset, p.Syntax, call.Pos(), "call_format") {
		return
	}

	if !IsViolation(p.Fset, call) {
		return
	}

	openLine := p.Fset.Position(call.Lparen).Line
	closeLine := p.Fset.Position(call.Rparen).Line

	if openLine == closeLine {
		results.AddConcern(
			concern.NewFile(
				"call_format",
				fmt.Sprintf(
					"call exceeds %d characters; split arguments to separate lines",
					maxLineLength,
				),
				p.Fset.Position(call.Pos()).Filename,
				false,
			),
		)
	} else {
		reportMultiLine(p, results, call)
	}
}
