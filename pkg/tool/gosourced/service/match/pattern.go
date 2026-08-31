package match

import "go/ast"

type Pattern struct {
	Holes     map[string]string
	Imports   map[string]string
	Statement ast.Stmt
}
