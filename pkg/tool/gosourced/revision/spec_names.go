package revision

import "go/ast"

func specNames(spec ast.Spec) []string {
	switch spec := spec.(type) {
	case *ast.TypeSpec:
		return []string{spec.Name.Name}
	case *ast.ValueSpec:
		var result []string

		for _, name := range spec.Names {
			result = append(result, name.Name)
		}

		return result
	default:
		return nil
	}
}
