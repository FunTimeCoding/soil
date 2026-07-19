package service

import (
	"go/ast"
	"go/types"
)

type moveEntry struct {
	symbol          string
	newName         string
	flipped         bool
	object          types.Object
	file            *ast.File
	declaration     ast.Decl
	spec            ast.Spec
	node            ast.Node
	carried         []*ast.ImportSpec
	targetFile      string
	backIdentifiers []*ast.Ident
}
