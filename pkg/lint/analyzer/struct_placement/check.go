package struct_placement

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func Check(
	p *packages.Package,
	results *output.Results,
) {
	scope := p.Types.Scope()
	var receivers []string
	var companions []*types.TypeName

	for _, name := range scope.Names() {
		t, okay := scope.Lookup(name).(*types.TypeName)

		if !okay || t.IsAlias() || generated(p, t.Pos()) {
			continue
		}

		named, okay := t.Type().(*types.Named)

		if !okay {
			continue
		}

		if _, isStruct := named.Underlying().(*types.Struct); !isStruct {
			continue
		}

		if named.NumMethods() > 0 {
			receivers = append(receivers, name)
		} else {
			companions = append(companions, t)
		}
	}

	if len(receivers) == 0 {
		return
	}

	owner := join.CommaSpace(receivers)

	for _, t := range companions {
		results.AddConcern(
			concern.NewPosition(
				"struct_placement",
				fmt.Sprintf(
					"struct %s shares the package with receiver struct %s - move %s to a bag or subpackage, or move %s out",
					t.Name(),
					owner,
					t.Name(),
					owner,
				),
				p.Fset.Position(t.Pos()),
			),
		)
	}
}
