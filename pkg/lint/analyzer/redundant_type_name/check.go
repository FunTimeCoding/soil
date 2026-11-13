package redundant_type_name

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/strings/camel"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func Check(
	p *packages.Package,
	results *output.Results,
) {
	scope := p.Types.Scope()
	identity := camel.FromSnake(p.Name)
	var structs []*types.TypeName

	for _, name := range scope.Names() {
		t, okay := scope.Lookup(name).(*types.TypeName)

		if !okay || t.IsAlias() || lint.IsGeneratedPosition(p, t.Pos()) {
			continue
		}

		named, okay := t.Type().(*types.Named)

		if !okay {
			continue
		}

		if named.NumMethods() > 0 || name == identity {
			return
		}

		if _, isStruct := named.Underlying().(*types.Struct); isStruct {
			structs = append(structs, t)
		}
	}

	if len(structs) < 2 {
		return
	}

	for _, t := range structs {
		stripped := Strip(t.Name(), identity)

		if stripped == "" {
			continue
		}

		message := fmt.Sprintf(
			"struct %s repeats the bag's name %s - rename it %s",
			t.Name(),
			p.Name,
			stripped,
		)

		if scope.Lookup(stripped) != nil {
			message = fmt.Sprintf(
				"struct %s repeats the bag's name %s, and %s is taken - name what tells %s from %s",
				t.Name(),
				p.Name,
				stripped,
				t.Name(),
				stripped,
			)
		}

		results.AddConcern(
			concern.NewPosition(
				"redundant_type_name",
				message,
				p.Fset.Position(t.Pos()),
			),
		)
	}
}
