package restricted_literal

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
					l, okay := n.(*ast.CompositeLit)

					if !okay {
						return true
					}

					checkLiteral(p, results, l, rule)

					return true
				},
			)
		}
	}
}
