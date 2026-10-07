package resolve_reference

import (
	"go/ast"
	"golang.org/x/tools/go/packages"
)

type Reference struct {
	Ident   *ast.Ident
	Package *packages.Package
}
