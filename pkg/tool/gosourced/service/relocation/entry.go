package relocation

import (
	"go/ast"
	"go/types"
)

type Entry struct {
	Symbol          string
	NewName         string
	Flipped         bool
	Object          types.Object
	File            *ast.File
	Declaration     ast.Decl
	Spec            ast.Spec
	Node            ast.Node
	Carried         []*ast.ImportSpec
	TargetFile      string
	BackIdentifiers []*ast.Ident
}
