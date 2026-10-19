package error_answer

import (
	"go/ast"
	"golang.org/x/tools/go/packages"
)

func listClasses(
	p *packages.Package,
	conditions []ast.Expr,
) []string {
	var result []string

	for _, c := range conditions {
		found := conditionClasses(p, c)

		if found == nil {
			return nil
		}

		result = append(result, found...)
	}

	return result
}
