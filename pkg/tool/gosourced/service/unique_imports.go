package service

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"go/ast"
)

func uniqueImports(entries []*moveEntry) []*ast.ImportSpec {
	seen := make(map[string]bool)
	var result []*ast.ImportSpec

	for _, entry := range entries {
		for _, spec := range entry.carried {
			key := spec.Path.Value

			if spec.Name != nil {
				key = join.Empty(spec.Name.Name, key)
			}

			if seen[key] {
				continue
			}

			seen[key] = true
			result = append(result, spec)
		}
	}

	return result
}
