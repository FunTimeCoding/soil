package symbol

import (
	"github.com/funtimecoding/soil/pkg/tool/gosourced/lookup"
	"go/types"
	"golang.org/x/tools/go/packages"
	"slices"
)

func methodObjects(
	loaded []*packages.Package,
	packagePath string,
	typeName string,
	methods []string,
) []types.Object {
	p := lookup.FindPackage(loaded, packagePath)

	if p == nil {
		return nil
	}

	n := namedType(p, typeName)

	if n == nil {
		return nil
	}

	var result []types.Object

	for m := range n.Methods() {
		if slices.Contains(methods, m.Name()) {
			result = append(result, m)
		}
	}

	return result
}
