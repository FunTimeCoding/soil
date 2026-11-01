package revision

import (
	"go/ast"
	"go/token"
	"path"
)

func packageDirectory(
	s *token.FileSet,
	f *ast.File,
) string {
	return path.Dir(s.File(f.Pos()).Name())
}
