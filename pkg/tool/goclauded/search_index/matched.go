package search_index

import (
	"fmt"
	"strings"
)

func matched(
	terms []string,
	kinds []string,
) (string, []any) {
	var parts []string
	var arguments []any

	for i, term := range terms {
		condition, argument := termCondition(term)
		parts = append(
			parts,
			fmt.Sprintf(
				`SELECT %d AS term, b.rowid AS row, b.session, b.at
				FROM block_text JOIN block b ON b.rowid = block_text.rowid
				WHERE %s AND b.kind IN (%s)`,
				i,
				condition,
				placeholders(len(kinds)),
			),
		)
		arguments = append(arguments, argument)

		for _, kind := range kinds {
			arguments = append(arguments, kind)
		}
	}

	return strings.Join(parts, " UNION ALL "), arguments
}
