package search_index

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/search_index/conversation"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/search_index/hit"
)

func (x *Index) matches(
	terms []string,
	kinds []string,
	limit int,
) ([]*conversation.Conversation, error) {
	union, arguments := matched(terms, kinds)
	rows, e := x.database.Query(
		fmt.Sprintf(
			`WITH matched AS MATERIALIZED (%s),
			complete AS (
				SELECT session, max(at) AS latest, count(DISTINCT row) AS hits
				FROM matched GROUP BY session
				HAVING count(DISTINCT term) = ?
				ORDER BY latest DESC LIMIT ?
			),
			scored AS (
				SELECT m.row, m.session, count(DISTINCT m.term) AS terms, m.at
				FROM matched m JOIN complete c ON c.session = m.session
				GROUP BY m.row
			),
			ranked AS (
				SELECT row, session, terms, at, row_number() OVER (
					PARTITION BY session ORDER BY terms DESC, at DESC
				) AS position
				FROM scored
			)
			SELECT c.session, c.latest, c.hits, b.identifier, b.role, b.kind, b.at, t.body
			FROM ranked r
			JOIN complete c ON c.session = r.session
			JOIN block b ON b.rowid = r.row
			JOIN block_text t ON t.rowid = r.row
			WHERE r.position <= ?
			ORDER BY c.latest DESC, r.session, r.position`,
			union,
		),
		append(arguments, len(terms), limit, constant.SearchSnippetLimit)...,
	)

	if e != nil {
		return nil, e
	}

	defer errors.PanicClose(rows)
	var result []*conversation.Conversation
	var current *conversation.Conversation

	for rows.Next() {
		var session, latest, identifier, role, kind, at, body string
		var count int

		if f := rows.Scan(
			&session,
			&latest,
			&count,
			&identifier,
			&role,
			&kind,
			&at,
			&body,
		); f != nil {
			return nil, f
		}

		if current == nil || current.Session != session {
			current = conversation.New(session, latest, count)
			result = append(result, current)
		}

		current.Hits = append(
			current.Hits,
			hit.New(identifier, role, kind, at, snippet(body, terms)),
		)
	}

	return result, rows.Err()
}
