package service

import "fmt"

func collectSpliceEdits(plan *movePlan) []*spliceEdit {
	var result []*spliceEdit

	for ident, name := range plan.renames {
		result = append(
			result,
			&spliceEdit{
				start:       ident.Pos(),
				end:         ident.End(),
				replacement: name,
			},
		)
	}

	for _, entry := range plan.entries {
		for _, ident := range entry.backIdentifiers {
			result = append(
				result,
				&spliceEdit{
					start: ident.Pos(),
					end:   ident.End(),
					replacement: fmt.Sprintf(
						"%s.%s",
						plan.sourceLocalName,
						ident.Name,
					),
				},
			)
		}
	}

	return result
}
