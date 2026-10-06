package service

import (
	"github.com/funtimecoding/soil/pkg/tool/gosourced/service/module_symbol"
	"go/types"
)

func moduleTargetPackages(
	directory string,
	symbols []*module_symbol.Symbol,
	modulePath string,
	newModulePath string,
) map[string]*types.Package {
	result := make(map[string]*types.Package)

	for _, symbol := range symbols {
		target := moduleTargetPath(
			symbol.PackagePath,
			modulePath,
			newModulePath,
		)

		if _, okay := result[target]; okay {
			continue
		}

		result[target] = loadModulePackage(directory, target)
	}

	return result
}
