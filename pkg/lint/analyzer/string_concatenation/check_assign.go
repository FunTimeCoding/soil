package string_concatenation

import (
	"github.com/funtimecoding/soil/pkg/lint/analyzer/suppress"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func checkAssign(
	p *packages.Package,
	results *output.Results,
	s *ast.AssignStmt,
) {
	t := p.TypesInfo.TypeOf(s.Lhs[0])

	if t == nil {
		return
	}

	b, okay := t.Underlying().(*types.Basic)

	if !okay || b.Kind() != types.String {
		return
	}

	if suppress.IsSuppressed(p.Fset, p.Syntax, s.Pos(), "string_concat") {
		return
	}

	results.AddConcern(
		concern.NewFile(
			"string_concatenation",
			"use fmt.Sprintf instead of string concatenation",
			p.Fset.Position(s.Pos()).Filename,
			false,
		),
	)
}
