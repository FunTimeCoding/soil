package service

import (
	"go/token"
	"go/types"
	"golang.org/x/tools/go/packages"
	"path/filepath"
)

func movedObjects(
	directory string,
	all []*packages.Package,
	set *token.FileSet,
	packagePath string,
	symbols []string,
	filePath string,
) []types.Object {
	p := findPackage(all, packagePath)

	if p == nil {
		return nil
	}

	if filePath != "" {
		full := filePath

		if !filepath.IsAbs(full) {
			full = filepath.Join(directory, filePath)
		}

		file := findSyntaxFile(set, p, full)

		if file == nil {
			return nil
		}

		expanded, e := expandFileSymbols(file)

		if e != nil {
			return nil
		}

		symbols = expanded
	}

	var result []types.Object

	for _, name := range symbols {
		if o := p.Types.Scope().Lookup(name); o != nil {
			result = append(result, o)
		}
	}

	return result
}
