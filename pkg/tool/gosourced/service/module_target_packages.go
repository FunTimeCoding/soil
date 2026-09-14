package service

import "go/types"

func moduleTargetPackages(
	directory string,
	symbols []*ModuleSymbol,
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
