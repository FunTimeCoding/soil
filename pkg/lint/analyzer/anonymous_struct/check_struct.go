package anonymous_struct

import (
	"github.com/funtimecoding/soil/pkg/lint/analyzer/suppress"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"go/ast"
	"golang.org/x/tools/go/packages"
)

func checkStruct(
	p *packages.Package,
	results *output.Results,
	s *ast.StructType,
) {
	if s.Fields == nil || len(s.Fields.List) == 0 {
		return
	}

	if isEmbeddingWrapper(s) {
		return
	}

	if suppress.IsSuppressed(p.Fset, p.Syntax, s.Pos(), "anonymous_struct") {
		return
	}

	results.AddConcern(
		concern.NewFile(
			"anonymous_struct",
			"anonymous struct; extract a named type",
			p.Fset.Position(s.Pos()).Filename,
			false,
		),
	)
}
