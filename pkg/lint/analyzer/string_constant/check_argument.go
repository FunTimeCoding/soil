package string_constant

import (
	"github.com/funtimecoding/soil/pkg/lint/analyzer/assert_call"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"go/ast"
	"go/token"
	"golang.org/x/tools/go/packages"
	"strings"
)

func checkArgument(
	p *packages.Package,
	results *output.Results,
	e ast.Expr,
	constants map[string][]knownConstant,
	expected []assert_call.Range,
) {
	l, okay := e.(*ast.BasicLit)

	if !okay || l.Kind != token.STRING {
		return
	}

	for _, r := range expected {
		if r.Contains(l.Pos()) {
			return
		}
	}

	value := strings.Trim(l.Value, "\"")

	if value == "" {
		return
	}

	list, exists := constants[value]

	if !exists {
		return
	}

	results.AddConcern(
		concern.NewFile(
			"string_constant",
			formatMessage(value, list),
			p.Fset.Position(l.Pos()).Filename,
			false,
		),
	)
}
