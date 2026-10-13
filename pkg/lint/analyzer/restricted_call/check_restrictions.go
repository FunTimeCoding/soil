package restricted_call

import (
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/lint/types/restriction"
	"go/ast"
	"golang.org/x/tools/go/packages"
	"slices"
)

func CheckRestrictions(
	p *packages.Package,
	results *output.Results,
	rules []restriction.Restriction,
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
					sel, okay := n.(*ast.SelectorExpr)

					if !okay {
						return true
					}

					checkSelector(p, results, sel, rule)

					return true
				},
			)
		}
	}
}
