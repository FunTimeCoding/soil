package imports

import (
	"go/ast"
	"go/token"
)

func positionInFile(
	fileSet *token.FileSet,
	file *ast.File,
	position token.Pos,
) bool {
	f := fileSet.File(file.Pos())

	return position >= file.Pos() && int(position) <= f.Base()+f.Size()
}
