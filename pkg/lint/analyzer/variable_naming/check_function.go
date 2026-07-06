package variable_naming

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"go/ast"
	"golang.org/x/tools/go/packages"
)

func checkFunction(
	p *packages.Package,
	results *output.Results,
	f *ast.FuncDecl,
) {
	variables := collectVariables(p.TypesInfo, f.Body)
	collectFromParams(p.TypesInfo, f, &variables)
	collectFromReceiver(p.TypesInfo, f, &variables)
	applyParameterExemptions(variables)

	if len(variables) == 0 {
		return
	}

	for i := range variables {
		variables[i].descendantNames = collectDescendantNames(
			p.TypesInfo,
			f,
			variables[i].ident,
		)
	}

	assignments := assignLetters(variables)

	for _, v := range variables {
		if !isEligible(v) {
			continue
		}

		ideal, okay := assignments[v.ident]

		if !okay {
			continue
		}

		if v.ident.Name == ideal {
			continue
		}

		results.AddConcern(
			concern.NewFile(
				"variable_naming",
				fmt.Sprintf(
					"variable %s of type %s should be named %s",
					v.ident.Name,
					v.typ.String(),
					ideal,
				),
				p.Fset.Position(v.ident.Pos()).Filename,
				false,
			),
		)
	}

	checkOkay(p, results, f)
}
