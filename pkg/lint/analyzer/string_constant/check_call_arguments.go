package string_constant

import (
	"github.com/funtimecoding/soil/pkg/lint/analyzer/assert_call"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"go/ast"
	"golang.org/x/tools/go/packages"
)

func checkCallArguments(
	p *packages.Package,
	results *output.Results,
	call *ast.CallExpr,
	constants map[string]knownConstant,
) {
	skipIndex := -1

	if assert_call.IsAssertCall(p, call) {
		skipIndex = 1
	}

	for i, a := range call.Args {
		if i == skipIndex {
			continue
		}

		checkArgument(p, results, a, constants)
	}
}
