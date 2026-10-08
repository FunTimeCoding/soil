package unexported_struct

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func Check(
	p *packages.Package,
	results *output.Results,
) {
	scope := p.Types.Scope()

	for _, name := range scope.Names() {
		t, okay := scope.Lookup(name).(*types.TypeName)

		if !okay || t.Exported() || t.IsAlias() ||
			lint.IsGeneratedPosition(p, t.Pos()) {
			continue
		}

		if _, isStruct := t.Type().Underlying().(*types.Struct); !isStruct {
			continue
		}

		results.AddConcern(
			concern.NewPosition(
				"unexported_struct",
				fmt.Sprintf("struct %s is unexported - export it", name),
				p.Fset.Position(t.Pos()),
			),
		)
	}
}
