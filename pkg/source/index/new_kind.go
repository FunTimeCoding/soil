package index

import (
	"go/types"
	"golang.org/x/tools/go/packages"
)

func NewKind(
	name string,
	value func() any,
	extract func(*packages.Package) any,
	external func(*types.Package) any,
) *Kind {
	return &Kind{
		Name:     name,
		Value:    value,
		Extract:  extract,
		External: external,
	}
}
