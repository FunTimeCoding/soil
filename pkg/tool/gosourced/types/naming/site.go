package naming

import (
	"go/ast"
	"golang.org/x/tools/go/packages"
)

type Site struct {
	Owner     *packages.Package
	File      *ast.File
	Signature *ast.FuncType
	Field     int
	Index     int
	Label     string
}
