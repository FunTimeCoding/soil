package constant_declaration

import (
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

// A named type is enum-shaped when its underlying type is a basic
// kind (string, int, float...) - the sanctioned residence for
// typed constant families inside constant/ homes.
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
