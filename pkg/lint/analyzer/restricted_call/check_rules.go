package restricted_call

import (
	"github.com/funtimecoding/soil/pkg/lint/output"
	"go/ast"
	"golang.org/x/tools/go/packages"
	"slices"
)

func CheckRules(
	p *packages.Package,
	results *output.Results,
	rules []Rule,
) {
	for _, rule := range rules {
		if slices.Contains(rule.AllowedIn, p.PkgPath) {
			continue
		}

		for _, file := range p.Syntax {
			if ast.IsGenerated(file) {
				continue
			}

			ast.Inspect(
				file,
				func(n ast.Node) bool {
					call, okay := n.(*ast.CallExpr)

					if !okay {
						return true
					}

					checkCall(p, results, call, rule)

					return true
				},
			)
		}
	}
}
