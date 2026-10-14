package unclosed_resource

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/lint/fact"
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func Extract(p *packages.Package) *fact.Summary {
	result := fact.NewSummary()

	if p.TypesInfo == nil {
		return result
	}

	for t, methods := range methodsByType(p) {
		if _, okay := methods[constant.CloseMethod]; okay {
			result.Releases[typeKey(t)] = releasedFields(p, methods)
		}
	}

	result.Shared = sharedVariables(p)

	for _, file := range p.Syntax {
		for _, d := range file.Decls {
			f, okay := d.(*ast.FuncDecl)

			if !okay || f.Body == nil {
				continue
			}

			o, okay := p.TypesInfo.Defs[f.Name].(*types.Func)

			if !okay {
				continue
			}

			k := functionKey(o)
			extractDisposals(p, f, o, result)
			extractTakes(p, f, o, result)
			extractWraps(p, f, o, result)

			if returnsSelf(p, f) {
				result.Self = append(result.Self, k)
			}

			var lending []string

			if lends(p, f, o, &lending) {
				if len(lending) > 0 {
					result.LendsVia[k] = lending
				} else {
					result.Lends = append(result.Lends, k)
				}
			}

			if spends(p, f, o) {
				if via := spentDelegates(p, f, o); len(via) > 0 {
					result.SpentVia[k] = via
				} else {
					result.Spent = append(result.Spent, k)
				}
			}

			conditions := registrarConditions(p, f.Body)

			if arranged(conditions) {
				result.Arranges = append(result.Arranges, k)

				continue
			}

			if len(conditions) > 0 {
				result.Arranging[k] = conditions
			}

			for _, g := range delegatesOf(p, f.Body) {
				if callee := functionKey(g); callee != k {
					result.Delegates[k] = append(result.Delegates[k], callee)
				}
			}
		}
	}

	return result
}
