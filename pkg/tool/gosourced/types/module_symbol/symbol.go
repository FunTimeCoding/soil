package module_symbol

import (
	"go/token"
	"go/types"
)

type Symbol struct {
	PackagePath string
	Owner       string
	Name        string
	Object      types.Object
	Position    token.Position
}
