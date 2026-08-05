package constant_declaration

import (
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func isEnumShaped(
	p *packages.Package,
	t *ast.TypeSpec,
) bool {
	definition, okay := p.TypesInfo.Defs[t.Name]

	if !okay {
		return false
	}

	_, okay = definition.Type().Underlying().(*types.Basic)

	return okay
}
