package defer_close

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func checkDefer(
	p *packages.Package,
	results *output.Results,
	d *ast.DeferStmt,
) {
	call, okay := d.Call.Fun.(*ast.SelectorExpr)

	if !okay || call.Sel.Name != "Close" {
		return
	}

	if len(d.Call.Args) != 0 {
		return
	}

	receiverType := p.TypesInfo.TypeOf(call.X)

	if receiverType == nil {
		return
	}

	if !implementsCloser(receiverType) {
		return
	}

	results.AddConcern(
		concern.NewFile(
			"defer_close",
			fmt.Sprintf(
				"use defer errors.PanicClose(%s) instead of defer %s.Close()",
				types.ExprString(call.X),
				types.ExprString(call.X),
			),
			p.Fset.Position(d.Pos()).Filename,
			false,
		),
	)
}
