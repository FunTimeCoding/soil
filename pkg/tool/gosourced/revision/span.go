package revision

import (
	"go/ast"
	"go/token"
)

func span(
	s *token.FileSet,
	n ast.Node,
) (int, int) {
	return s.Position(n.Pos()).Offset, s.Position(n.End()).Offset
}
