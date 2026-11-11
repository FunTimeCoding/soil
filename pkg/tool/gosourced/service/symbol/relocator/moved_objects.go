package relocator

import (
	"github.com/funtimecoding/soil/pkg/tool/gosourced/lookup"
	"go/token"
	"go/types"
	"golang.org/x/tools/go/packages"
	"path/filepath"
)

func MovedObjects(
	directory string,
	all []*packages.Package,
	set *token.FileSet,
	packagePath string,
	symbols []string,
	filePath string,
) []types.Object {
	p := lookup.FindPackage(all, packagePath)

	if p == nil {
		return nil
	}

	if filePath != "" {
		full := filePath

		if !filepath.IsAbs(full) {
			full = filepath.Join(directory, filePath)
		}

		file := lookup.FindSyntaxFile(set, p, full)

		if file == nil {
			return nil
		}

		expanded, e := ExpandFileSymbols(file)

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
