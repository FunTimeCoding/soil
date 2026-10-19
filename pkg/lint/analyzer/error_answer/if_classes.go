package error_answer

import (
	"go/ast"
	"golang.org/x/tools/go/packages"
)

func ifClasses(
	p *packages.Package,
	x *ast.IfStmt,
) []string {
	if c := asTypeClass(p, x); c != "" {
		return []string{c}
	}

	return conditionClasses(p, x.Cond)
}
