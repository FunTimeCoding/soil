package service

import (
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func findDeclaration(
	all []*packages.Package,
	packagePath string,
	symbol string,
	receiver string,
) (types.Object, *packages.Package, error) {
	var target *packages.Package

	for _, p := range all {
		if p.PkgPath == packagePath {
			target = p

			break
		}
	}

	if target == nil {
		return nil, nil, not_found.New("package", packagePath)
	}

	if receiver == "" {
		return findFunction(target, symbol)
	}

	return findMethod(target, symbol, receiver)
}
