package stray_variable

import (
	"go/ast"
	"golang.org/x/tools/go/packages"
	"strings"
)

func isSynchronization(
	p *packages.Package,
	name *ast.Ident,
) bool {
	o := p.TypesInfo.Defs[name]

	if o == nil {
		return false
	}

	kind := strings.TrimPrefix(o.Type().String(), "*")

	return strings.HasPrefix(kind, "sync.") ||
		strings.HasPrefix(kind, "sync/atomic.")
}
