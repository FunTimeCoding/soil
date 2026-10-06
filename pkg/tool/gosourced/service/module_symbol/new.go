package module_symbol

import (
	"go/token"
	"go/types"
)

func New(
	packagePath string,
	owner string,
	name string,
	object types.Object,
	position token.Position,
) *Symbol {
	return &Symbol{
		PackagePath: packagePath,
		Owner:       owner,
		Name:        name,
		Object:      object,
		Position:    position,
	}
}
