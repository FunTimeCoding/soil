package unclosed_resource

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"github.com/funtimecoding/soil/pkg/lint/fact"
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func closeEquivalents(
	p *packages.Package,
	t *types.TypeName,
	methods map[string]*ast.FuncDecl,
	summary *fact.Summary,
) {
	required := closableReleases(
		t,
		releasedFrom(p, methods, constant.CloseMethod),
	)

	if len(required) == 0 {
		return
	}

	for name, m := range methods {
		if name == constant.CloseMethod {
			continue
		}

		o, okay := p.TypesInfo.Defs[m.Name].(*types.Func)

		if !okay {
			continue
		}

		if conditions, covers := coverage(
			required,
			releasedFrom(p, methods, name),
		); covers {
			k := functionKey(o)
			summary.Forwards[k] = append(
				summary.Forwards[k],
				fact.Forward{
					Argument:   constant.ReceiverArgument,
					Conditions: conditions,
				},
			)
		}
	}
}
