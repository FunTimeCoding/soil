package search_index

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"strings"
)

func (x *Index) MatchingBlocks(
	session string,
	query string,
	kinds []string,
) (map[string]bool, error) {
	result := map[string]bool{}
	kinds = kindsOrDefault(kinds)

	for _, term := range strings.Fields(query) {
		condition, argument := termCondition(term)
		arguments := []any{argument, session}

		for _, kind := range kinds {
			arguments = append(arguments, kind)
		}

		rows, e := x.database.Query(
			fmt.Sprintf(
				`SELECT b.identifier
				FROM block_text JOIN block b ON b.rowid = block_text.rowid
				WHERE %s AND b.session = ? AND b.kind IN (%s)`,
				condition,
				placeholders(len(kinds)),
			),
			arguments...,
		)

		if e != nil {
			return nil, e
		}

		for rows.Next() {
			var identifier string

			if f := rows.Scan(&identifier); f != nil {
				errors.PanicClose(rows)

				return nil, f
			}

			result[identifier] = true
		}

		errors.PanicClose(rows)

		if f := rows.Err(); f != nil {
			return nil, f
		}
	}

	return result, nil
}
