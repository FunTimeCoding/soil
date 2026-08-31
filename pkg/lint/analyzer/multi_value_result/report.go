package multi_value_result

import (
	"github.com/funtimecoding/soil/pkg/lint/analyzer/suppress"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"go/ast"
	"golang.org/x/tools/go/packages"
)

func report(
	p *packages.Package,
	results *output.Results,
	call *ast.CallExpr,
	message string,
) {
	if suppress.IsSuppressed(
		p.Fset,
		p.Syntax,
		call.Pos(),
		"multi_value_result",
	) {
		return
	}

	results.AddConcern(
		concern.NewFile(
			"multi_value_result",
			message,
			p.Fset.Position(call.Pos()).Filename,
			false,
		),
	)
}
