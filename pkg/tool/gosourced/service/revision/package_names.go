package revision

import (
	"go/ast"
	"go/token"
)

func packageNames(
	s *token.FileSet,
	parsed []*ast.File,
) map[string]map[string]bool {
	result := map[string]map[string]bool{}

	for _, f := range parsed {
		key := packageKey(packageDirectory(s, f), f.Name.Name)

		if result[key] == nil {
			result[key] = map[string]bool{}
		}

		for _, name := range declaredNames(f) {
			result[key][name] = true
		}
	}

	return result
}
