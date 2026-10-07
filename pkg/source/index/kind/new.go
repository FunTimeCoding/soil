package kind

import (
	"go/types"
	"golang.org/x/tools/go/packages"
)

func New(
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
