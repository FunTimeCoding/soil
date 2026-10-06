package service

import (
	"go/ast"
	"go/token"
	"go/types"
	"golang.org/x/tools/go/ast/astutil"
)

func definitionPath(
	file *ast.File,
	o types.Object,
) []ast.Node {
	if file == nil || o.Pos() < file.Pos() || o.Pos() > file.End() {
		return nil
	}

	path, _ := astutil.PathEnclosingInterval(
		file,
		o.Pos(),
		o.Pos()+token.Pos(len(o.Name())),
	)

	if len(path) == 0 {
		return nil
	}

	if i, okay := path[0].(*ast.Ident); !okay || i.Pos() != o.Pos() {
		return nil
	}

	return path
}
