package variable_naming

import (
	"go/ast"
	"golang.org/x/tools/go/packages"
)

func ComputeRenames(
	p *packages.Package,
	u *ast.FuncDecl,
) []Rename {
	variables := collectVariables(p.TypesInfo, u.Body)
	collectFromParams(p.TypesInfo, u, &variables)
	collectFromReceiver(p.TypesInfo, u, &variables)
	applyParameterExemptions(variables)
	var result []Rename

	if len(variables) > 0 {
		for i := range variables {
			variables[i].descendantNames = collectDescendantNames(
				p.TypesInfo,
				u,
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

			o := p.TypesInfo.ObjectOf(v.ident)

			if o == nil {
				continue
			}

			result = append(
				result,
				Rename{
					Object:  o,
					NewName: ideal,
				},
			)
		}
	}

	forms := collectOkayForms(p.TypesInfo, u)

	if len(forms) > 0 {
		okayAssignments := assignOkayNames(forms)

		for _, f := range forms {
			expected := okayAssignments[f.ident]

			if f.ident.Name == expected {
				continue
			}

			o := p.TypesInfo.ObjectOf(f.ident)

			if o == nil {
				continue
			}

			result = append(
				result,
				Rename{
					Object:  o,
					NewName: expected,
				},
			)
		}
	}

	return result
}
