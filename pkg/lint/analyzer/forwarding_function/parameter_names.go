package forwarding_function

import "go/ast"

func parameterNames(f *ast.FuncDecl) []string {
	if f.Type.Params == nil {
		return nil
	}

	var result []string

	for _, field := range f.Type.Params.List {
		if len(field.Names) == 0 {
			return nil
		}

		for _, name := range field.Names {
			if name.Name == "_" {
				return nil
			}

			result = append(result, name.Name)
		}
	}

	return result
}
