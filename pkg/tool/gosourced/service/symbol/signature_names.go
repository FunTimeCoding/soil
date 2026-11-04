package symbol

import "go/ast"

func signatureNames(t *ast.FuncType) map[string]bool {
	result := map[string]bool{}

	for _, list := range []*ast.FieldList{t.Params, t.Results} {
		if list == nil {
			continue
		}

		for _, field := range list.List {
			for _, n := range field.Names {
				if n.Name != "_" {
					result[n.Name] = true
				}
			}
		}
	}

	return result
}
