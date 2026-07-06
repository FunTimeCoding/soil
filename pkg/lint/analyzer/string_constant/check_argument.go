package string_constant

import (
	"fmt"
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
	constants map[string]knownConstant,
) {
	l, okay := e.(*ast.BasicLit)

	if !okay || l.Kind != token.STRING {
		return
	}

	value := strings.Trim(l.Value, "\"")

	if value == "" {
		return
	}

	c, exists := constants[value]

	if !exists {
		return
	}

	results.AddConcern(
		concern.NewFile(
			"string_constant",
			fmt.Sprintf(
				"string literal %q has constant %s.%s",
				value,
				c.packageName,
				c.name,
			),
			p.Fset.Position(l.Pos()).Filename,
			false,
		),
	)
}
