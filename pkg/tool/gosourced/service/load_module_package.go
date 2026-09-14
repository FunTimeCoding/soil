package service

import (
	"github.com/funtimecoding/soil/pkg/source/resolve"
	"go/types"
)

func loadModulePackage(
	directory string,
	packagePath string,
) *types.Package {
	all, _, e := resolve.LoadPackages(directory, packagePath)

	if e != nil {
		return nil
	}

	for _, loaded := range all {
		if loaded.PkgPath == packagePath && loaded.Types != nil {
			return loaded.Types
		}
	}

	return nil
}
