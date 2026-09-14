package service

import (
	"go/token"
	"go/types"
)

type ModuleSymbol struct {
	PackagePath string
	Owner       string
	Name        string
	Object      types.Object
	Position    token.Position
}
