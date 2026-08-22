package error_wrap_verb

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint/analyzer/suppress"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func checkCall(
	p *packages.Package,
	results *output.Results,
	call *ast.CallExpr,
) {
	selector, okay := call.Fun.(*ast.SelectorExpr)

	if !okay {
		return
	}

	o, okay := p.TypesInfo.Uses[selector.Sel]

	if !okay {
		return
	}

	f, okay := o.(*types.Func)

	if !okay {
		return
	}

	a := f.Pkg()

	if a == nil || a.Path() != "fmt" || f.Name() != "Errorf" {
		return
	}

	if len(call.Args) < 2 || call.Ellipsis.IsValid() {
		return
	}

	format, okay := formatValue(p, call.Args[0])

	if !okay {
		return
	}

	verbs, okay := parseVerbs(format)

	if !okay {
		return
	}

	for index, argument := range call.Args[1:] {
		letter, present := verbs[index]

		if !present || (letter != 's' && letter != 'v') {
			continue
		}

		if !isErrorType(p.TypesInfo.TypeOf(argument)) {
			continue
		}

		if suppress.IsSuppressed(
			p.Fset,
			p.Syntax,
			call.Pos(),
			"error_wrap_verb",
		) {
			return
		}

		results.AddConcern(
			concern.NewFile(
				"error_wrap_verb",
				fmt.Sprintf(
					"error argument formatted with %%%c - use %%w so errors.Is and errors.As unwrap the chain",
					letter,
				),
				p.Fset.Position(call.Pos()).Filename,
				false,
			),
		)

		return
	}
}
