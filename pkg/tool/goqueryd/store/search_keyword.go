package store

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"math"
)

func (s *Store) SearchKeyword(
	query string,
	limit int,
	collection string,
	full bool,
) ([]SearchResult, error) {
	fullTextSearch := BuildFullTextSearchQuery(query)

	if fullTextSearch == "" {
		return nil, nil
	}

	sql := `
		WITH fts_matches AS (
			SELECT rowid, bm25(document_full_text_search, 1.5, 4.0, 1.0) AS bm25_score
			FROM document_full_text_search
			WHERE document_full_text_search MATCH ?
			ORDER BY bm25_score ASC
			LIMIT ?
		)
		SELECT
			d.collection || '/' || d.path AS filepath,
			d.collection,
			d.path,
			d.title,
			d.hash,
			content.body,
			fm.bm25_score
		FROM fts_matches fm
		JOIN document d ON d.identifier = fm.rowid
		JOIN content ON content.hash = d.hash
		WHERE d.active = 1`
	arguments := []any{fullTextSearch, limit * 10}

	if collection != "" {
		sql = join.Empty(sql, " AND d.collection = ?")
		arguments = append(arguments, collection)
	}

	sql = join.Empty(sql, " ORDER BY fm.bm25_score ASC LIMIT ?")
	arguments = append(arguments, limit)
	rows, e := s.database.Query(sql, arguments...)

	if e != nil {
		return nil, e
	}

	defer errors.PanicClose(rows)
	var result []SearchResult

	for rows.Next() {
		var r SearchResult
		var bm25 float64
		var body string

		if f := rows.Scan(
			&r.FilePath,
			&r.Collection,
			&r.Path,
			&r.Title,
			&r.Hash,
			&body,
			&bm25,
		); f != nil {
			return nil, f
		}

		r.VirtualPath = buildVirtualPath(r.Collection, r.Path)
		r.Context = s.ResolveContext(r.Collection, r.Path)
		r.Score = math.Abs(bm25) / (1 + math.Abs(bm25))
		r.Source = "full_text_search"
		snippet, line := ExtractSnippet(body, query, 0)
		r.Snippet = snippet
		r.SnippetLine = line

		if full {
			r.Body = body
		}

		result = append(result, r)
	}

	if e := rows.Err(); e != nil {
		return nil, e
	}

	return result, nil
}
