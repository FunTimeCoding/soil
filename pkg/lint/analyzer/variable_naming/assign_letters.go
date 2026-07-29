package variable_naming

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"go/ast"
	"sort"
)

func assignLetters(variables []typedVariable) map[*ast.Ident]string {
	result := make(map[*ast.Ident]string)
	taken := make(map[string]bool)
	singleLetterVars := filterEligible(variables)

	for _, v := range singleLetterVars {
		taken[v.ident.Name] = true
	}

	sort.SliceStable(
		singleLetterVars,
		func(i, j int) bool {
			pi := scopePriority(singleLetterVars[i].kind)
			pj := scopePriority(singleLetterVars[j].kind)

			if pi != pj {
				return pi < pj
			}

			if singleLetterVars[i].precedence != singleLetterVars[j].precedence {
				return singleLetterVars[i].precedence < singleLetterVars[j].precedence
			}

			return singleLetterVars[i].ident.Pos() < singleLetterVars[j].ident.Pos()
		},
	)
	taken = make(map[string]bool)
	localNames := map[string]bool{}

	for _, v := range singleLetterVars {
		if v.kind == constant.VariableKindLocal {
			localNames[v.ident.Name] = true
		}
	}

	for _, v := range singleLetterVars {
		perVariableTaken := make(
			map[string]bool,
			len(taken)+len(v.scopedNames)+len(v.descendantNames),
		)

		for k := range taken {
			perVariableTaken[k] = true
		}

		if v.kind == constant.VariableKindLocal {
			for k := range v.scopedNames {
				perVariableTaken[k] = true
			}

			for k := range v.descendantNames {
				if k == v.ident.Name {
					continue
				}

				perVariableTaken[k] = true
			}
		} else {
			for k := range v.scopedNames {
				if localNames[k] {
					continue
				}

				perVariableTaken[k] = true
			}
		}

		letter := pickLetter(v.typ, perVariableTaken)

		if letter == "" {
			continue
		}

		result[v.ident] = letter
		taken[letter] = true
	}

	return result
}
