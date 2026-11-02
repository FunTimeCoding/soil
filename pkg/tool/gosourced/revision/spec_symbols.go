package revision

import (
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/symbol"
	"go/ast"
	"go/token"
)

func specSymbols(
	s *token.FileSet,
	f *ast.File,
	v *ast.ValueSpec,
	kind string,
	body string,
	shape string,
) []*symbol.Symbol {
	position := s.Position(v.Pos())
	var result []*symbol.Symbol

	for _, name := range v.Names {
		entry := symbol.New(
			packageDirectory(s, f),
			f.Name.Name,
			"",
			name.Name,
			kind,
			position.Filename,
			position.Line,
		)
		entry.Body = body
		entry.Shape = shape
		result = append(result, entry)
	}

	return result
}
